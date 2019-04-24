package table

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	table := New(9, 9)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i >= j {
				table.Cell(i-1, j-1).SetString(fmt.Sprintf("%d*%d=%d", i, j, i*j))
			}
		}
	}

	table.Style.SetBorderColor(Red)

	table.Print()
}
