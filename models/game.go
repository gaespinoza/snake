package models

import "errors"

type Game struct {
	Board *Board
	Snake *List
	Score int
	Timer int
}

func NewGame(rows, columns int) (*Game, error) {
	board := NewBoard(rows, columns)
	if board == nil {
		return nil, errors.New("invalid board dimensions")
	}
	snake := NewSnake()
	board.Cells[0][0].Filled = true
	return &Game{
		Board: board,
		Snake: snake,
		Score: 0,
		Timer: 0,
	}, nil
}
