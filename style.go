package table

type style struct {
	fontColor    color
	fontBgColor  color
	fontWeight   weight
	borderColor  color
	borderWeight weight
	decorations  []decoration
}

type color int

const (
	Black color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type weight int

const (
	Bolder weight = iota + 1
	Lighter
)

type decoration int

const (
	Oblique decoration = iota + 3
	Underline
	Flicker
	_
	Reverse
	Blanking
)

// functions of style

func (s *style) SetFontColor(c color) {
	s.fontColor = c
}

func (s *style) SetFontBgColor(c color) {
	s.fontBgColor = c
}

func (s *style) SetFontWeight(w weight) {
	s.fontWeight = w
}

func (s *style) SetBorderColor(c color) {
	s.borderColor = c
}

func (s *style) SetBorderWeight(w weight) {
	s.borderWeight = w
}

func (s *style) SetDecoration(decorations ...decoration) {
	s.decorations = decorations
}
