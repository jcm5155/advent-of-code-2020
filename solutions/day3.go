package solutions

import (
	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day3 solution
func (h *Handler) Day3() (int, int) {
	pzl := util.ReadPuzzleInput("3")
	return d3(pzl)
}

func d3(pzl []string) (int, int) {
	tree := byte('#')
	yMax, xMax := len(pzl), len(pzl[0])
	slopes, xPos, treeCounts := [5]int{1, 3, 5, 7, 1}, [5]int{}, [5]int{}

	for y := 0; y < yMax; y++ {
		for i, x := range xPos {
			if i == 4 && y%2 != 0 {
				continue
			}
			if pzl[y][x] == tree {
				treeCounts[i]++
			}
			xPos[i] = (x + slopes[i]) % xMax
		}
	}

	p1 := treeCounts[1]
	p2 := treeCounts[0]
	for i := 1; i < len(treeCounts); i++ {
		p2 *= treeCounts[i]
	}
	return p1, p2
}
