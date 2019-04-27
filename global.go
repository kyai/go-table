package table

var (
	asciiSize = 2
	emojiSize = 1
)

func SetAsciiSize(i int) {
	asciiSize = i
}

func SetEmojiSize(i int) {
	emojiSize = i
}

func size(s string) int {
	size := 0
	for _, r := range s {
		l := len(string(r))
		switch l {
		case 1:
			size += 1
		case 3:
			size += asciiSize
		case 4:
			size += emojiSize
		default:
			size += l
		}
	}
	return size
}

func max(i ...int) int {
	max := i[0]
	for _, v := range i {
		if v > max {
			max = v
		}
	}
	return max
}

func min(i ...int) int {
	min := i[0]
	for _, v := range i {
		if v < min {
			min = v
		}
	}
	return min
}
