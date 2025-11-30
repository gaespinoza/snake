package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gaespinoza/snake/models"
	"github.com/gaespinoza/snake/state"
)

func main() {
	window := new(app.Window)
	var ops op.Ops
	ui := state.NewMainState()
	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))
	go func() {
		changes := time.NewTicker(time.Second / 2)
		defer changes.Stop()
		for _ = range changes.C {
			window.Invalidate()
		}
	}()
	go func() {
		err := run(window, ui, &ops, theme)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window, ui *state.MainUi, ops *op.Ops, theme *material.Theme) error {
	for {
		switch e := window.Event().(type) {

		case key.Event:
			// Only process navigation keys if we are in the game state
			if ui.CurrentState == state.GameState && ui.Game != nil && e.State == key.Press {
				switch e.Name {
				case key.NameUpArrow:
					ui.Game.Model.Snake.SetDirection(models.Up)
				case key.NameDownArrow:
					ui.Game.Model.Snake.SetDirection(models.Down)
				case key.NameLeftArrow:
					ui.Game.Model.Snake.SetDirection(models.Left)
				case key.NameRightArrow:
					ui.Game.Model.Snake.SetDirection(models.Right)
				default:
					log.Printf("unhandled key: %s", e.Name)
				}
			}

		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(ops, e)

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
				// Pressed a key?
				ev, ok := gtx.Event(

					key.Filter{Optional: key.ModShift, Name: key.NameUpArrow},
					key.Filter{Optional: key.ModShift, Name: key.NameDownArrow},
					key.Filter{Optional: key.ModShift, Name: key.NameLeftArrow},
					key.Filter{Optional: key.ModShift, Name: key.NameRightArrow},
				)
				if ok {
					switch ev.(key.Event).Name {
					case key.NameUpArrow:
						ui.Game.Model.Snake.SetDirection(models.Up)
					case key.NameDownArrow:
						ui.Game.Model.Snake.SetDirection(models.Down)
					case key.NameLeftArrow:
						ui.Game.Model.Snake.SetDirection(models.Left)
					case key.NameRightArrow:
						ui.Game.Model.Snake.SetDirection(models.Right)
					}
				}
				if ui.Game.BackButton.Clicked(gtx) {
					ui.CurrentState = state.HomeState
					ui.Game = nil
				}

				// 2. Render the Board
				// We draw the board every frame, regardless of whether the snake moved
				state.GetGameLayout(gtx, theme, ui.Game)

				// 3. Handle Game Logic (Time Gated)
				// Only move if enough time has passed since the last move
				now := time.Now()
				if ui.Game != nil && now.Sub(ui.Game.Model.LastMove) >= ui.Game.Model.Speed {

					ui.Game.Model.MoveSnake()
					log.Printf("Snake Head Direction: %v", ui.Game.Model.Snake.Direction)

					// if err != nil {
					// 	log.Printf("Game Over: %v", err)
					// 	ui.CurrentState = state.HomeState
					// 	ui.Game = nil
					// 	// Important: Return/Break here to prevent accessing nil ui.Game below
					// 	break
					// }

					// Update timestamp only on successful move
					ui.Game.Model.LastMove = now

					// Force a redraw immediately after a logic update so the UI reflects the move
					window.Invalidate()
				}

			default:
				log.Fatalf("unknown state: %s", ui.CurrentState)
			}

			e.Frame(gtx.Ops)

		}

	}
}
