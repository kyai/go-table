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
	layout layout
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

func (t *table) Row(n int) *cellset {
	return &cellset{
		cells: t.cells[n],
	}
}

func (t *table) Col(n int) *cellset {
	cells := make([]*cell, 0)
	for r := 0; r < t.Rows; r++ {
		cells = append(cells, t.cells[r][n])
	}
	return &cellset{
		cells: cells,
	}
}

type layout int

const (
	Average layout = iota
	Adaptive
)

func (t *table) SetLayout(l layout) {
	t.layout = l
}

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
	// Width of each column
	widthPreCol := make([]int, t.Cols)

	if t.thead != nil {
		for k, th := range t.thead.cells {
			widthPreCol[k] = size(th.Text)
		}
	}

	for r := 0; r < t.Rows; r++ {
		for c := 0; c < t.Cols; c++ {
			if ce := t.cells[r][c]; ce != nil {
				if l := size(ce.Text); l > widthPreCol[c] {
					widthPreCol[c] = l
				}
			}
		}
	}

	var widths []int

	if t.layout == Adaptive {
		for i := 0; i < t.Cols; i++ {
			widths = append(widths, widthPreCol[i]+blankSize*2)
		}
	} else {
		maxWidth := max(widthPreCol...)
		for i := 0; i < t.Cols; i++ {
			widths = append(widths, maxWidth+blankSize*2)
		}
	}

	t.widths = widths
}
