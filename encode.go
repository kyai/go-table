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

func (t *table) encodeThead() (s string) {
	if t.thead == nil {
		return
	}

	vertical := t.symbol.Vertical
	vertical = parseColor(vertical, t.Style.borderColor)

	for c := 0; c < t.Cols; c++ {
		ce := &cell{}
		if c < len(t.thead.cells) {
			ce = t.thead.cells[c]
		}
		s += vertical
		s += t.encodeCell(ce, t.widths[c], t.thead.Style)
	}
	s += vertical
	s += "\n"
	s += t.encodeMiddle()
	return
}

func (t *table) encodeCell(c *cell, width int, s *style) string {
	text := c.Text
	l := size(text)
	if l > width {
		return ""
	}

	if (width-l)%2 > 0 {
		text += blank
		l += 1
	}

	bl := (width - l) / 2
	for i := 0; i < bl; i++ {
		text = blank + text + blank
	}

	if fontColor := lastValue(s.fontColor, c.Style.fontColor); fontColor != nil {
		text = parseColor(text, fontColor.(color))
	}
	if fontBgColor := lastValue(s.fontBgColor, c.Style.fontBgColor); fontBgColor != nil {
		text = parseBgColor(text, fontBgColor.(color))
	}
	if fontWeight := lastValue(s.fontWeight, c.Style.fontWeight); fontWeight != nil {
		text = parseWeight(text, fontWeight.(weight))
	}

	return text
}

func lastValue(vs ...interface{}) interface{} {
	var r interface{}
	for _, v := range vs {
		n := 0
		switch v.(type) {
		case color:
			n = int(v.(color))
		case weight:
			n = int(v.(weight))
		}
		if n != 0 {
			r = v
		}
	}
	return r
}

func (t *table) encodeCells() (s string) {
	vertical := t.symbol.Vertical
	vertical = parseColor(vertical, t.Style.borderColor)

	for r := 0; r < t.Rows; r++ {
		for c := 0; c < t.Cols; c++ {
			s += vertical
			s += t.encodeCell(t.cells[r][c], t.widths[c], t.Style)
		}
		s += vertical
		s += "\n"
		if r < t.Rows-1 {
			s += t.encodeMiddle()
		}
	}
	return
}
