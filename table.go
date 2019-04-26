package table

import (
	"fmt"
)

type table struct {
	Rows   int
	Cols   int
	cells  [][]*cell
	thead  *thead
	Style  *style
	symbol *Symbol
	widths []int
}

type thead struct {
	cells []*cell
	Style *style
}

func New(rows, cols int) *table {
	var cells [][]*cell
	for r := 0; r < rows; r++ {
		slice := make([]*cell, cols)
		for c := 0; c < cols; c++ {
			slice[c] = &cell{
				Row:   r,
				Col:   c,
				Style: &style{},
			}
		}
		cells = append(cells, slice)
	}

	return &table{
		Rows:   rows,
		Cols:   cols,
		cells:  cells,
		Style:  &style{},
		symbol: DefaultSymbol,
	}
}

// TODO: row+1
func NewThead(tds ...string) *thead {
	var cells []*cell
	for k, td := range tds {
		cells = append(cells, &cell{
			Row:   0,
			Col:   k,
			Text:  td,
			Style: &style{},
		})
	}

	return &thead{
		cells: cells,
		Style: &style{},
	}
}

func (t *table) SetThead(th *thead) {
	t.thead = th
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
	out += t.encodeThead()
	out += t.encodeCells()
	out += t.encodeTopOrBottom(false)
	return
}

func (t *table) varWidths() {
	maxWidth := 0

	if t.thead != nil {
		for _, th := range t.thead.cells {
			if l := size(th.Text); l > maxWidth {
				maxWidth = l
			}
		}
	}

	for r := 0; r < t.Rows; r++ {
		for c := 0; c < t.Cols; c++ {
			if ce := t.cells[r][c]; ce != nil {
				if l := size(t.cells[r][c].Text); l > maxWidth {
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
