package main

import (
	"github.com/GodsBoss/ld45"
	"github.com/GodsBoss/ld45/pkg/console"
	"github.com/GodsBoss/ld45/pkg/date"
	"github.com/GodsBoss/ld45/pkg/dom"
	"github.com/GodsBoss/ld45/pkg/listeners"
	"github.com/GodsBoss/ld45/pkg/loop"
	"github.com/GodsBoss/ld45/ui"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	listeners.Add(js.Global.Get("window"), "load", initialize)
}

func initialize(this *js.Object, arguments []*js.Object) interface{} {
	doc, err := dom.GlobalDocument()
	if err != nil {
		console.Log("Initialization failed", err.Error())
		return nil
	}
	canvas := doc.CreateCanvas(800, 600)
	doc.GetElementByID("game").AppendChild(canvas)
	dom.SetStyles(
		canvas,
		map[string]string{
			"border":       "1px solid #000",
			"display":      "block",
			"margin-left":  "auto",
			"margin-right": "auto",
			"width":        "800px",
		},
	)
	img := doc.CreateImage()
	img.OnLoad(runGame(canvas, img))
	img.OnError(
		func(event listeners.Event) {
			console.Log("error", event.Type())
		},
	)
	img.Src("gfx.png")
	return nil
}

func runGame(canvas *dom.Canvas, img *dom.Image) func() {
	return func() {
		tps := 50
		game := ld45.NewGame()
		timed := &loop.Timed{
			CurrentTimeInMS: func() int {
				return date.Now().Unix()
			},
			ScheduleStep: func(f func(), delay int) {
				dom.GlobalWindow().SetTimeout(f, delay)
			},
		}
		timed.Start(
			func() {
				game.Tick(1000 / tps)
			},
			1000/tps,
		)

		ctx := canvas.GetContext2D()
		ctx.DisableImageSmoothing()
		renderer := &ui.Renderer{
			Zoom:          2,
			Ctx:           ctx,
			ImageSource:   dom.ImageElementSource(img),
			SpriteMapping: createSpriteMapping(),
			BackgroundMapping: map[string]dom.FillStyle{
				"title":            dom.Color("#334422"),
				"playing":          dom.Color("#334422"),
				"game_over":        dom.Color("#334422"),
				"choose_character": dom.Color("#334422"),
			},
		}
		simple := &loop.Simple{
			ScheduleStep: func(f func()) {
				dom.GlobalWindow().RequestAnimationFrame(
					func(_ float64) {
						f()
					},
				)
			},
		}
		simple.Start(
			func() {
				renderer.Draw(game)
			},
		)
		keyEventMapping := map[string]ld45.KeyEventType{
			"keyup":    ld45.KeyUp,
			"keydown":  ld45.KeyDown,
			"keypress": ld45.KeyPress,
		}
		passKeyToGame := func(event listeners.KeyEvent) {
			if event.Repeat() {
				return
			}
			eventType, ok := keyEventMapping[event.Type()]
			if !ok {
				return
			}
			game.InvokeKeyEvent(
				ld45.KeyEvent{
					Type:  eventType,
					Alt:   event.AltKey(),
					Ctrl:  event.CtrlKey(),
					Shift: event.ShiftKey(),
					Key:   event.Key(),
				},
			)
		}
		dom.GlobalWindow().
			OnKeyDown(passKeyToGame).
			OnKeyUp(passKeyToGame).
			OnKeyPress(passKeyToGame)
	}
}
