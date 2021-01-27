package web

import (
	"image"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/gurupras/dota2bot"
	"github.com/gurupras/dota2bot/gui"
	"github.com/sirupsen/logrus"
)

type WebUI struct {
	*gui.BaseGUI
	*http.ServeMux
	updateChan  chan struct{}
	connections map[*websocket.Conn]struct{}
}

func NewWebUI(gameInfo *dota2bot.GameInfo) *WebUI {
	mux := http.NewServeMux()
	rect := image.Rect(0, 0, 1024, 1024)
	baseGUI := gui.NewBaseGUI(gameInfo, rect)
	ui := &WebUI{
		baseGUI,
		mux,
		make(chan struct{}),
		make(map[*websocket.Conn]struct{}),
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/home/guru/workspace/go/src/github.com/gurupras/dota2bot/gui/web/static/"))))
	mux.HandleFunc("/ws", ui.ws)
	mux.HandleFunc("/", ui.root)
	return ui
}

func (w *WebUI) root(res http.ResponseWriter, req *http.Request) {
	logrus.Infof("Entered serveHTTP")
	http.ServeFile(res, req, "/home/guru/workspace/go/src/github.com/gurupras/dota2bot/gui/web/static/index.html")
}

func (w *WebUI) Update(units []dota2bot.Unit) {
	w.BaseGUI.Update(units)
	w.updateChan <- struct{}{}
}

func (w *WebUI) ws(res http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		logrus.Errorf("upgrade error: %v", err)
		return
	}
	w.connections[c] = struct{}{}
	for {
		defer func() {
			delete(w.connections, c)
		}()
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (w *WebUI) Run() error {
	logrus.Infof("Server can be accessed from http://localhost:4105/")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _ = range w.updateChan {
			for k, _ := range w.connections {
				k.WriteJSON(w.LastUpdate)
			}
		}
	}()
	http.ListenAndServe(":4105", w.ServeMux)
	wg.Wait()
	return nil
}
