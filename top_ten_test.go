package topten

import (
	"testing"
)

func TestTopTen(t *testing.T) {
	got := TopTen(`
		1
		2 2
		3 3 3
		4 4 4 4
		5 5 5 5 5
		6 6 6 6 6 6
		7 7 7 7 7 7 7
		8 8 8 8 8 8 8 8
		9 9 9 9 9 9 9 9 9
		10 10 10 10 10 10 10 10 10 10
		11 11 11 11 11 11 11 11 11 11 11
		12 12 12 12 12 12 12 12 12 12 12 12
		13 13 13 13 13 13 13 13 13 13 13 13 13
		14 14 14 14 14 14 14 14 14 14 14 14 14 14
		15 15 15 15 15 15 15 15 15 15 15 15 15 15 15
	`)
	reference := [...]string{
		"15", "14", "13", "12", "11", "10", "9", "8", "7", "6"}

	for i, _ := range make([]int, 10) {
		if got[i] != reference[i] {
			t.Errorf("get[%d] = %s; want %s", i, got[i], reference[i])
		}
	}
}
