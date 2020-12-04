package solutions

import (
	"strconv"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day1 solution
func (h *Handler) Day1() (int, int) {
	p := util.ReadPuzzleInput("1", "\n")
	pzl := []int{}
	for _, i := range p {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		pzl = append(pzl, j)
	}
	return d1p1(pzl), d1p2(pzl)
}

func d1p1(pzl []int) int {
	for i, ival := range pzl {
		c := 2020 - ival
		for j := i + 1; j < len(pzl); j++ {
			if pzl[j] == c {
				return ival * pzl[j]
			}
		}
	}
	panic("WHOOPSIES - Didn't find an answer for Day 1 Part 1")
}

func d1p2(pzl []int) int {
	for i := range pzl {
		for j := i + 1; j < len(pzl); j++ {
			if pzl[i]+pzl[j] < 2020 {
				for k := j + 1; k < len(pzl); k++ {
					if pzl[i]+pzl[j]+pzl[k] == 2020 {
						return pzl[i] * pzl[j] * pzl[k]
					}
				}
			}
		}
	}
	panic("WHOOPSIES - Didn't find an answer for Day 1 Part 2")
}
