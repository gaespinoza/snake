package main

import (
	"log"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gaespinoza/snake/state"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	ui := state.NewMainState()
	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))

	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			switch ui.CurrentState {
			case state.HomeState:
				// handle button click logic
				if ui.Home.StartButton.Clicked(gtx) {
					rows, re := strconv.Atoi(ui.Home.HeightInput.Text())
					columns, ce := strconv.Atoi(ui.Home.WidthInput.Text())
					if re != nil || ce != nil {
						log.Printf("invalid size input: %v", e)
						break
					}

					stateGame, err := state.NewGameState(rows, columns)
					if err != nil {
						log.Printf("failed to create game state: %v", err)
						break
					}
					ui.Game = stateGame
					ui.CurrentState = state.GameState
				}

				state.GetHomeStateLayout(gtx, theme, ui.Home)
				// Pass the drawing operations to the GPU
			case state.GameState:
				// Check for Back Button Click
				if ui.Game.BackButton.Clicked(gtx) {
					ui.CurrentState = state.HomeState
					// ui.Game = nil // Reset game state
				}
				state.GetGameLayout(gtx, theme, ui.Game)
			default:
				log.Fatalf("unknown state: %s", ui.CurrentState)
			}

			e.Frame(gtx.Ops)

		}
	}
}
