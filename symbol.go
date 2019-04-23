package table

type Symbol struct {
	Horizontal  string
	Vertical    string
	Left        string
	Right       string
	Top         string
	Bottom      string
	LeftTop     string
	RightTop    string
	LeftBottom  string
	RightBottom string
	Middle      string
}

var DefaultSymbol *Symbol

func init() {
	DefaultSymbol = &Symbol{"─", "│", "├", "┤", "┬", "┴", "┌", "┐", "└", "┘", "┼"}
}
