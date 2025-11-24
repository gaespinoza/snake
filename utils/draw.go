package utils

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Helper function to draw a labeled input field
func DrawInput(gtx layout.Context, th *material.Theme, editor *widget.Editor, label string) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			l := material.Body2(th, label)
			l.Color = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
			return l.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			e := material.Editor(th, editor, "20")

			// Style the editor
			border := widget.Border{
				Color:        color.NRGBA{A: 255, R: 200, G: 200, B: 200},
				CornerRadius: unit.Dp(4),
				Width:        unit.Dp(2),
			}

			// Add padding inside the border
			return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, e.Layout)
			})
		}),
	)
}
