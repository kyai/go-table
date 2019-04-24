package table

type style struct {
	fontColor    color
	fontBgColor  color
	fontWeight   weight
	borderColor  color
	borderWeight weight
	isOblique    bool
	isUnderline  bool
}

type color int

const (
	Black color = iota + 30
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
	Normal weight = iota
	Bolder
	Lighter
)

type decoration int

const (
	None decoration = iota
	Oblique
	Underline
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
	for _, v := range decorations {
		switch v {
		case None:
			s.isOblique = false
			s.isUnderline = false
			break
		case Oblique:
			s.isOblique = true
		case Underline:
			s.isUnderline = true
		}
	}
}
