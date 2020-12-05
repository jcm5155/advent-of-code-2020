package solutions

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day5 solution
func (h *Handler) Day5() (int, int) {
	r := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	pzl := strings.Split(r.Replace(util.ReadRawPuzzleInput("5")), "\n")
	seatIDs := []int{}
	p1, p2 := 0, 0
	for _, s := range pzl {
		row, _ := strconv.ParseInt(s[:7], 2, 64)
		col, _ := strconv.ParseInt(s[7:], 2, 64)
		seatID := int(row*8 + col)
		seatIDs = append(seatIDs, seatID)
		if seatID > p1 {
			p1 = seatID
		}
	}
	sort.Ints(seatIDs)
	for i, v := range seatIDs {
		if seatIDs[i+1]-v == 2 {
			p2 = v + 1
			break
		}
	}
	return p1, p2
}
