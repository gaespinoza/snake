package models

type CellState string

const (
	EmptyCell  CellState = "empty"
	FilledCell CellState = "filled"
	FoodCell   CellState = "food"
	HeadCell   CellState = "head"
)

type Cell struct {
	Row    int
	Column int
	State  CellState
}

type Board struct {
	Rows    int
	Columns int
	Cells   [][]Cell
}

func NewBoard(rows, columns int) *Board {
	if rows <= 0 || columns <= 0 {
		return nil
	}
	cells := make([][]Cell, rows)
	for r := range rows {
		cells[r] = make([]Cell, columns)
		for c := range columns {
			cells[r][c] = Cell{Row: r, Column: c, State: EmptyCell}
		}
	}
	return &Board{
		Rows:    rows,
		Columns: columns,
		Cells:   cells,
	}
}

func (b *Board) CellHasFood(row, column int) bool {
	if row < 0 || row >= b.Rows || column < 0 || column >= b.Columns {
		return false
	}
	return b.Cells[row][column].State == FoodCell
}

func (b *Board) PlaceFood(row, column int) bool {
	if row < 0 || row >= b.Rows || column < 0 || column >= b.Columns {
		return false
	}
	b.Cells[row][column].State = FoodCell
	return true
}

func (b *Board) RemoveFood(row, column int) bool {
	if row < 0 || row >= b.Rows || column < 0 || column >= b.Columns {
		return false
	}
	b.Cells[row][column].State = EmptyCell
	return true
}

func (b *Board) CellIsFilled(row, column int) bool {
	if row < 0 || row >= b.Rows || column < 0 || column >= b.Columns {
		return false
	}
	return b.Cells[row][column].State == FilledCell
}

func (b *Board) SetCellFilled(row, column int) bool {
	if row < 0 || row >= b.Rows || column < 0 || column >= b.Columns {
		return false
	}
	b.Cells[row][column].State = FilledCell
	return true
}

func (b *Board) SetCellEmpty(row, column int) bool {
	if row < 0 || row >= b.Rows || column < 0 || column >= b.Columns {
		return false
	}
	b.Cells[row][column].State = EmptyCell
	return true
}

func (b *Board) SetCellHead(row, column int) bool {
	if row < 0 || row >= b.Rows || column < 0 || column >= b.Columns {
		return false
	}
	b.Cells[row][column].State = HeadCell
	return true
}
