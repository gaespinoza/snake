package state

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/gaespinoza/snake/utils"
)

type HomeUi struct {
	WidthInput  widget.Editor
	HeightInput widget.Editor
	PlayerName  widget.Editor
	StartButton widget.Clickable

	IsGameRunning bool
}

func NewHomeState() *HomeUi {
	home := &HomeUi{
		WidthInput:    widget.Editor{SingleLine: true, Submit: true},
		HeightInput:   widget.Editor{SingleLine: true, Submit: true},
		PlayerName:    widget.Editor{SingleLine: true, Submit: true},
		StartButton:   widget.Clickable{},
		IsGameRunning: false,
	}

	home.WidthInput.SetText("20")
	home.HeightInput.SetText("20")
	home.PlayerName.SetText("Player1")
	return home
}

// layoutLandingPage defines the visual structure of the home screen
func GetHomeStateLayout(gtx layout.Context, th *material.Theme, ui *HomeUi) layout.Dimensions {
	// Use a Vertical Flex layout to stack elements
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle, // Center horizontally
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx,
		// --- Title ---
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Center the title specifically
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				h1 := material.H3(th, "SNAKE")
				h1.Color = color.NRGBA{R: 0, G: 100, B: 0, A: 255} // Dark Green
				h1.Font.Weight = font.Bold
				return h1.Layout(gtx)
			})
		}),

		// --- Spacing ---
		layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),

		// --- Width Input ---
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return utils.DrawInput(gtx, th, &ui.WidthInput, "Board Width")
		}),

		// --- Spacing ---
		layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),

		// --- Height Input ---
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return utils.DrawInput(gtx, th, &ui.HeightInput, "Board Height")
		}),

		// --- Spacing ---
		layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),

		// --- Height Input ---
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return utils.DrawInput(gtx, th, &ui.PlayerName, "Player Name")
		}),

		// --- Spacing ---
		layout.Rigid(layout.Spacer{Height: unit.Dp(30)}.Layout),

		// --- Start Button ---
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Add margins around the button
			margins := layout.Inset{
				Left:  unit.Dp(50),
				Right: unit.Dp(50),
			}
			return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &ui.StartButton, "START GAME")
				btn.Background = color.NRGBA{R: 46, G: 125, B: 50, A: 255} // Green button
				return btn.Layout(gtx)
			})
		}),
	)
}
