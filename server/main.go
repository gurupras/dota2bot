package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/gurupras/dota2bot"
	"github.com/gurupras/dota2bot/commands"
	"github.com/gurupras/dota2bot/gui"
	"github.com/gurupras/dota2bot/gui/web"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

var gameInfo *dota2bot.GameInfo
var radiant = dota2bot.NewTeam("Radiant")
var dire = dota2bot.NewTeam("Dire")
var radiantHandler httpHandler
var direHandler httpHandler

var miniMap gui.GUI

var mutex sync.Mutex
var wg sync.WaitGroup

// create a handler struct
type httpHandler struct {
	*dota2bot.Team
	*http.ServeMux
	worldUpdateIndex int
}

func newHTTPHandler(team *dota2bot.Team) httpHandler {
	mux := http.NewServeMux()
	h := httpHandler{
		team,
		mux,
		0,
	}
	mux.HandleFunc("/", h.log)

	mux.HandleFunc("/ws", h.ws)
	return h
}

func parseMessage(b []byte, format string, data interface{}) error {
	switch format {
	case "J":
		if err := json.Unmarshal(b, data); err != nil {
			return err
		}
	case "M":
		arrays := bytes.Split(b, []byte{'!', '!', '@'})
		b = bytes.Join(arrays, []byte{0})
		if err := msgpack.Unmarshal(b, data); err != nil {
			return err
		}
	}
	return nil
}

func (h httpHandler) log(res http.ResponseWriter, req *http.Request) {
	team := h.Team
	response := make([]commands.Command, 0)
	b, _ := ioutil.ReadAll(req.Body)
	// logrus.Debugf("Bytes: \n%v\nstr:\n%v\n", b, string(b))
	// First byte determines JSON/msgpack
	format := string(b[0:1])
	// Next 3 bytes is length of messageType
	messageTypeLenBytes := b[1:4]
	messageTypeLen, _ := strconv.Atoi(string(messageTypeLenBytes))
	messageType := string(b[4 : 4+messageTypeLen])
	dataBytes := b[4+messageTypeLen:]
	// logrus.Debugf("[%v]: format=%v messageType=%v", h.name, format, messageType)
	switch messageType {
	case "gameinfo":
		gi := dota2bot.GameInfo{}
		if err := parseMessage(dataBytes, format, &gi); err != nil {
			logrus.Warnf("[%v]: Failed to decode message: %v\n%v\n", h.Name, string(b), err)
		}
		gameInfo = &gi
		team.TeamID = gi.TeamID
		mutex.Lock()
		if miniMap == nil {
			miniMap = web.NewWebUI(gameInfo)
			wg.Add(1)
			go func() {
				defer wg.Done()
				miniMap.Run()
			}()
		}
		mutex.Unlock()
		logrus.Infof("[%v]: Parsed gameinfo", h.Name)
	case "observation":
		obs := dota2bot.Unit{}
		if err := parseMessage(dataBytes, format, &obs); err != nil {
			logrus.Warnf("[%v]: Failed to decode message: %v\n%v\n", h.Name, string(b), err)
		}
		if gameInfo == nil {
			return
		}
		team.Lock()
		defer team.Unlock()
		oldObs := team.Bots[obs.PlayerID]
		if oldObs.ActiveMode != obs.ActiveMode {
			logrus.Infof("[%v]: %v switched mode: %v", h.Name, obs.Name, gameInfo.GetBotModeString(obs.ActiveMode))
		}
		team.Bots[obs.PlayerID] = obs
		if team.DriverID == 0 || team.DriverID == obs.PlayerID {
			team.DriverID = obs.PlayerID
			response = append(response, commands.CommandDesignateToSendWorldInfo())
			logrus.Debugf("[%v]: Designating bot as worldinfo driver: %v (%v)", h.Name, obs.PlayerID, obs.Name)
		}
		// logrus.Debugf("Observation by: %v", obs.PlayerID)
	case "world":
		wi := dota2bot.WorldInfo{}
		if err := parseMessage(dataBytes, format, &wi); err != nil {
			logrus.Warnf("[%v]: Failed to decode message: %v", h.Name, err)
		}
		logrus.Debugf("[%v]: World info: units=%v", h.Name, len(wi.Units))
		team.WorldInfo = wi
		h.worldUpdateIndex++
		mutex.Lock()
		if miniMap != nil && radiantHandler.worldUpdateIndex == direHandler.worldUpdateIndex {
			logrus.Debugf("[%v]: Updating minimap: team=%v", h.Name, team.TeamID)
			miniMap.Update(append(radiant.Units, dire.Units...))
		}
		mutex.Unlock()
	case "log":
		data := make(map[string]string)
		if err := parseMessage(dataBytes, format, &data); err != nil {
			logrus.Warnf("[%v]: Failed to decode message: %v", h.Name, err)
		}
		switch data["level"] {
		case "debug":
			logrus.Debugln(data["message"])
		case "info":
			logrus.Infoln(data["message"])
		case "warn":
			logrus.Warnln(data["message"])
		case "error":
			logrus.Errorln(data["message"])
		}
	}
	// logrus.Infof("[%v]: [%v] - %v\n", h.Name, messageType, data)
	responseBytes, _ := json.Marshal(response)
	res.Write(append([]byte("J000"), responseBytes...))
}

func (h httpHandler) ws(res http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		logrus.Printf("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			logrus.Println("read:", err)
			break
		}
		logrus.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			logrus.Println("write:", err)
			break
		}
	}
}

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	// create a new handler
	radiantHandler = newHTTPHandler(radiant)
	direHandler = newHTTPHandler(dire)

	// listen and serve
	wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		logrus.Infof("Starting %v", radiantHandler.Name)
		http.ListenAndServe(":4101", radiantHandler)
	}()
	go func() {
		defer wg.Done()
		logrus.Infof("Starting %v", direHandler.Name)
		http.ListenAndServe(":4102", direHandler)
	}()

	wg.Wait()

}
