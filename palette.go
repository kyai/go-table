package table

import "fmt"

func parseColor(s string, c color) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}
