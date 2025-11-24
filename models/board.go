package models

type Cell struct {
	Row    int
	Column int
	Filled bool
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
			cells[r][c] = Cell{Row: r, Column: c, Filled: false}
		}
	}
	return &Board{
		Rows:    rows,
		Columns: columns,
		Cells:   cells,
	}
}
