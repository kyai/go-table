package table

import (
	"strconv"
)

type cell struct {
	Text  string
	Row   int
	Col   int
	Style *style
}

type cellset struct {
	cells []*cell
}

func (t *table) Cell(row, col int) *cell {
	return t.cells[row][col]
}

func (c *cell) set(text string) {
	c.Text = text
}

// method of cell

func (c *cell) SetString(v string) {
	c.set(v)
}

func (c *cell) SetInt(v int) {
	c.set(strconv.Itoa(v))
}

// method of cellset

func (cs *cellset) SetString(v string) {
	for _, c := range cs.cells {
		c.SetString(v)
	}
}

func (cs *cellset) SetInt(v int) {
	for _, c := range cs.cells {
		c.SetInt(v)
	}
}

func (cs *cellset) SetStyle(s *style) {
	for _, c := range cs.cells {
		c.Style = s
	}
}
