package main

import (
	"os"
	"time"

	gt "github.com/kyai/go-table"
)

func main() {
	t := time.Now()
	if args := os.Args; len(args) > 1 {
		t, _ = time.Parse("200601", args[1])
	}
	t = t.AddDate(0, 0, -t.Day()+1)
	t = t.AddDate(0, 0, -int(t.Weekday()))

	ts := make([]time.Time, 42)
	for k, _ := range ts {
		ts[k] = t
		t = t.AddDate(0, 0, 1)
	}

	table := gt.New(6, 7)

	thead := gt.NewThead("Sun", "Mon", "Tues", "Wed", "Thur", "Fri", "Sat")
	thead.Style.SetFontWeight(gt.Bolder)
	table.SetThead(thead)

	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			d := ts[row*7+col]
			c := table.Cell(row, col)
			c.SetInt(d.Day())

			if d.Format("060102") == time.Now().Format("060102") {
				c.Style.SetFontBgColor(gt.Red)
			}
		}
	}

	table.Print()
}
