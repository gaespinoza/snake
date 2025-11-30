package state

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gaespinoza/snake/models"
)

var (
	Green     color.NRGBA = color.NRGBA{R: 0, G: 100, B: 0, A: 255}
	DarkGreen color.NRGBA = color.NRGBA{R: 0, G: 150, B: 0, A: 255}
	White     color.NRGBA = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	LightGray color.NRGBA = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	Red       color.NRGBA = color.NRGBA{R: 150, G: 0, B: 0, A: 255}
)

type GameUi struct {
	BackButton widget.Clickable

	Model *models.Game
}

func NewGameState(rows, columns int) (*GameUi, error) {
	game, err := models.NewGame(rows, columns)
	if err != nil {
		return nil, err
	}
	return &GameUi{
		Model: game,
	}, nil
}

func GetGameLayout(gtx layout.Context, th *material.Theme, game *GameUi) layout.Dimensions {
	if game == nil {
		return layout.Dimensions{}
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// 1. Header (Back Button + Score)
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &game.BackButton, "BACK")
						btn.Background = Red
						btn.Inset = layout.UniformInset(unit.Dp(8))
						return btn.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.H6(th, fmt.Sprintf("Score: %d", game.Model.Score))
						return lbl.Layout(gtx)
					}),
				)
			})
		}),

		// 2. The Game Board Area
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// We use a custom widget here to draw the board
			return layoutBoard(gtx, game.Model)
		}),
	)
}

// layoutBoard handles the custom drawing of the grid and snake
func layoutBoard(gtx layout.Context, model *models.Game) layout.Dimensions {
	// 1. Calculate Cell Size
	// We want the board to fit within the available space while maintaining aspect ratio
	availSize := gtx.Constraints.Max
	rows := model.Board.Rows
	cols := model.Board.Columns

	// Safe check to avoid divide by zero
	if rows == 0 || cols == 0 {
		return layout.Dimensions{Size: availSize}
	}

	cellW := availSize.X / cols
	cellH := availSize.Y / rows

	// Use the smaller dimension to keep cells square
	cellSize := cellW
	if cellH < cellW {
		cellSize = cellH
	}

	// 2. Center the board in the available space
	boardWidth := cols * cellSize
	boardHeight := rows * cellSize
	offsetX := (availSize.X - boardWidth) / 2
	offsetY := (availSize.Y - boardHeight) / 2

	// Wrap the drawing operations in an Offset to center it
	defer op.Offset(image.Point{
		X: offsetX,
		Y: offsetY,
	}).Push(gtx.Ops).Pop()

	// 3. Draw Background (The Grid)
	// Light grey background for the whole board
	paint.FillShape(gtx.Ops,
		LightGray,
		clip.Rect{Max: image.Pt(boardWidth, boardHeight)}.Op(),
	)

	// Draw grid lines (optional, but helps visualization)
	// We iterate every cell to draw borders or backgrounds
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Calculate position
			x := c * cellSize
			y := r * cellSize

			// Draw a faint border for each cell
			// We do this by drawing a slightly smaller rect inside the cell area
			// or just by having a background color and drawing cells with gaps.
			// Here is a simple stroke effect by drawing a smaller rect:
			padding := 2 // 1 pixel gap
			cellRect := image.Rect(x+padding, y+padding, x+cellSize-padding, y+cellSize-padding)

			colorToUse := LightGray

			switch model.Board.Cells[r][c].State {
			case models.FilledCell:
				colorToUse = Green
			case models.HeadCell:
				colorToUse = DarkGreen
			case models.FoodCell:
				colorToUse = Red
			case models.EmptyCell:
			default:
				colorToUse = White
			}

			paint.FillShape(gtx.Ops,
				colorToUse,
				clip.Rect(cellRect).Op(),
			)
		}
	}

	return layout.Dimensions{Size: availSize}
}
