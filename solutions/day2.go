package solutions

import (
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day2 solution
func (h *Handler) Day2() (int, int) {
	pzl := util.ReadPuzzleInput("2")
	return d2p1(pzl), d2p2(pzl)
}

func d2Process(i string) (int, int, byte, string) {
	parts := strings.Split(i, " ")
	lims := strings.Split(parts[0], "-")
	minVal, _ := strconv.Atoi(lims[0])
	maxVal, _ := strconv.Atoi(lims[1])
	val := parts[1][0]
	pass := parts[2]
	return minVal, maxVal, val, pass
}

func d2p1(pzl []string) int {
	validPassCount := 0
	for _, i := range pzl {
		minVal, maxVal, val, pass := d2Process(i)
		valCount := 0
		for _, c := range []byte(pass) {
			if c == val {
				valCount++
			}
		}
		if minVal <= valCount && maxVal >= valCount {
			validPassCount++
		}
	}
	return validPassCount
}

func d2p2(pzl []string) int {
	validPassCount := 0
	for _, i := range pzl {
		minVal, maxVal, val, pass := d2Process(i)
		minVal--
		maxVal--
		exactlyOne := false

		if minVal <= len(pass) && pass[minVal] == val {
			exactlyOne = !exactlyOne
		}

		if maxVal <= len(pass) && pass[maxVal] == val {
			exactlyOne = !exactlyOne
		}

		if exactlyOne == true {
			validPassCount++
		}
	}
	return validPassCount
}
