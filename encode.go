package table

var (
	blank     = " "
	blankSize = 1
)

func (t *table) encodeTopOrBottom(isTop bool) (s string) {
	if isTop {
		s += t.symbol.LeftTop
	} else {
		s += t.symbol.LeftBottom
	}

	for k, width := range t.widths {
		if k > 0 {
			if isTop {
				s += t.symbol.Top
			} else {
				s += t.symbol.Bottom
			}
		}
		for i := 0; i < width; i++ {
			s += t.symbol.Horizontal
		}
	}

	if isTop {
		s += t.symbol.RightTop
	} else {
		s += t.symbol.RightBottom
	}

	s += "\n"
	s = parseColor(s, t.Style.borderColor)
	return
}

func (t *table) encodeMiddle() (s string) {
	s += t.symbol.Left
	for k, width := range t.widths {
		if k > 0 {
			s += t.symbol.Middle
		}
		for i := 0; i < width; i++ {
			s += t.symbol.Horizontal
		}
	}
	s += t.symbol.Right
	s += "\n"
	s = parseColor(s, t.Style.borderColor)
	return
}

func (t *table) encodeCell(text string, width int) (s string) {
	l := len(text)
	if l > width {
		return
	}

	if (width-l)%2 > 0 {
		text += blank
		l += 1
	}

	bl := (width - l) / 2
	for i := 0; i < bl; i++ {
		text = blank + text + blank
	}

	return text
}

func (t *table) encodeCells() (s string) {
	vertical := t.symbol.Vertical
	vertical = parseColor(vertical, t.Style.borderColor)

	for r := 0; r < t.Rows; r++ {
		for c := 0; c < t.Cols; c++ {
			s += vertical
			s += t.encodeCell(t.cells[r][c].Text, t.widths[c])
		}
		s += vertical
		s += "\n"
		if r < t.Rows-1 {
			s += t.encodeMiddle()
		}
	}
	return
}
