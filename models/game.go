package models

import (
	"errors"
	"time"
)

type Game struct {
	Board *Board
	Snake *Snake
	Score int

	// Logic Control
	LastMove time.Time     // When did we last move?
	Speed    time.Duration // How often should we move?
}

func NewGame(rows, columns int) (*Game, error) {
	board := NewBoard(rows, columns)
	if board == nil {
		return nil, errors.New("invalid board dimensions")
	}
	snake := NewSnake()
	board.SetCellHead(0, 0)

	return &Game{
		Board:    board,
		Snake:    snake,
		Score:    0,
		LastMove: time.Now(),
		Speed:    time.Second / 4, // Move 4 times per second
	}, nil
}

func (g *Game) MoveSnake() error {

	oldHead := g.Snake.Head

	g.Board.SetCellFilled(oldHead.Row, oldHead.Column)
	g.Snake.AddToHead()

	if g.SnakeHeadOutOfBounds() {
		return errors.New("snake moved out of bounds")
	}

	head := g.Snake.Head

	if g.Board.CellIsFilled(head.Row, head.Column) {
		return errors.New("snake collided with itself")
	}

	g.Board.SetCellHead(head.Row, head.Column)

	if !g.Board.CellHasFood(head.Row, head.Column) {
		tail := g.Snake.Tail
		g.Board.SetCellEmpty(tail.Row, tail.Column)
		g.Snake.RemoveFromTail()
	} else {
		g.Board.RemoveFood(head.Row, head.Column)
		g.Score += 10
	}
	return nil
}

func (g *Game) SnakeHeadOutOfBounds() bool {
	head := g.Snake.Head
	if head.Row < 0 || head.Row >= g.Board.Rows || head.Column < 0 || head.Column >= g.Board.Columns {
		return true
	}
	return false
}
