package table

import "fmt"

func printer(s string, i int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", i, s)
}

func parseColor(s string, c color) string {
	if c == 0 {
		return s
	}
	return printer(s, int(c)+30)
}

func parseBgColor(s string, c color) string {
	if c == 0 {
		return s
	}
	return printer(s, int(c)+40)
}

func parseWeight(s string, w weight) string {
	return printer(s, int(w))
}
