package main

import (
	"github.com/golang.org/x/mobile/app"
	"github.com/golang.org/x/mobile/event/lifecycle"
	"github.com/golang.org/x/mobile/event/paint"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := app.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					a.Send(paint.Event{})
				}
			case paint.Event:
				a.Publish()
				a.Send(paint.Event{}) // keep animating
			}
		}
	})
}
