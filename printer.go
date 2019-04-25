package table

import "fmt"

func printer(s string, i int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", i, s)
}

func parseColor(s string, c color) string {
	return printer(s, int(c))
}
