package solutions

import (
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day8 solution
func (h *Handler) Day8() (int, int) {
	pzl := util.ReadPuzzleInput("8", "\n")
	c := make(chan int, 1)
	p1 := d8resolveTimeline(pzl, c)

	for idx, val := range pzl {
		select {
		case p2 := <-c:
			return p1, p2
		default:
			break
		}
		cmd := val[:3]
		switch cmd {
		case "nop":
			cmd = "jmp"
		case "jmp":
			cmd = "nop"
		default:
			continue
		}
		var cp = make([]string, len(pzl))
		copy(cp, pzl)
		cp[idx] = cmd + val[3:]
		go d8resolveTimeline(cp, c)
	}
	// block just in case
	p2 := <-c
	return p1, p2
}

func d8resolveTimeline(pzl []string, c chan int) int {
	idx, acc := 0, 0
	idxVisited := map[int]bool{}
	for {
		if idx == len(pzl)-1 {
			c <- acc
			return 0
		}
		if idxVisited[idx] {
			return acc
		}

		idxVisited[idx] = true
		parts := strings.Fields(pzl[idx])
		cmd := parts[0]
		amt, _ := strconv.Atoi(parts[1])

		switch cmd {
		case "acc":
			acc += amt
			idx++
		case "nop":
			idx++
		case "jmp":
			idx += amt
		default:
			panic("this should never happen")
		}
	}
}
