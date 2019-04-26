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

func (t *table) Cell(row, col int) *cell {
	return t.cells[row][col]
}

func (c *cell) set(text string) {
	c.Text = text
}

func (c *cell) SetString(s string) {
	c.set(s)
}

func (c *cell) SetInt(i int) {
	c.set(strconv.Itoa(i))
}
