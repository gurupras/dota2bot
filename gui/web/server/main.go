package main

import (
	"time"

	"github.com/gurupras/dota2bot/gui"
	"github.com/gurupras/dota2bot/gui/web"
)

func main() {
	gi := gui.FakeGameInfo()
	w := web.NewWebUI(gi)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			units := gui.FakeUpdate()
			w.Update(units)
		}
	}()
	w.Run()
}
