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
	ids := []int{}
	for _, s := range pzl {
		row, _ := strconv.ParseInt(s[:7], 2, 64)
		col, _ := strconv.ParseInt(s[7:], 2, 64)
		ids = append(ids, int(row*8+col))
	}
	sort.Ints(ids)
	for i, v := range ids {
		if ids[i+1]-v == 2 {
			return ids[len(ids)-1], v + 1
		}
	}
	panic("well this is awkward")
}
