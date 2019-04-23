package table

import (
	"fmt"
)

type table struct {
	Rows   int
	Cols   int
	cells  [][]*cell
	symbol *Symbol
	widths []int
}

func New(rows, cols int) *table {
	var cells [][]*cell
	for r := 0; r < rows; r++ {
		slice := make([]*cell, cols)
		for c := 0; c < cols; c++ {
			slice[c] = &cell{}
		}
		cells = append(cells, slice)
	}

	return &table{
		Rows:   rows,
		Cols:   cols,
		cells:  cells,
		symbol: DefaultSymbol,
	}
}

func (t *table) Row() {}

func (t *table) Col() {}

func (t *table) SetSymbol(s *Symbol) {
	t.symbol = s
}

func (t *table) Print() {
	fmt.Println(t.toString())
}

func (t *table) toString() (out string) {
	t.varWidths()
	out += t.encodeTopOrBottom(true)
	out += t.encodeCells()
	out += t.encodeTopOrBottom(false)
	return
}

func (t *table) varWidths() {
	maxWidth := 0
	for r := 0; r < t.Rows; r++ {
		for c := 0; c < t.Cols; c++ {
			if ce := t.cells[r][c]; ce != nil {
				if l := len(t.cells[r][c].Text); l > maxWidth {
					maxWidth = l
				}
			}
		}
	}

	maxWidth += blankSize * 2

	var widths []int
	for i := 0; i < t.Cols; i++ {
		widths = append(widths, maxWidth)
	}

	t.widths = widths
}
