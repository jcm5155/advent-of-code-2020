package solutions

import (
	"strings"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day6 solution
func (h *Handler) Day6() (int, int) {
	pzl := util.ReadPuzzleInput("6", "\n\n")
	p1, p2 := 0, 0
	for _, group := range pzl {
		people := strings.Fields(group)
		numPeople := len(people)
		joinedPeople := strings.Join(people, "")
		m := make(map[rune]int)
		for _, c := range joinedPeople {
			m[c]++
		}
		for _, v := range m {
			p1++
			if v == numPeople {
				p2++
			}
		}
	}
	return p1, p2
}
