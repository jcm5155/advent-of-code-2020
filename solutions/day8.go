package solutions

import (
	"strconv"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day8 solution
func (h *Handler) Day8() (int, int) {
	pzl := util.ReadPuzzleInput("8", "\n")
	c := make(chan int, 1)
	p1 := d8resolveTimeline(pzl, -1, "", c)

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
		go d8resolveTimeline(pzl, idx, cmd+val[3:], c)
	}
	// block just in case
	p2 := <-c
	return p1, p2
}

func d8resolveTimeline(pzl []string, replaceIdx int, replaceVal string, c chan int) int {
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
		val := pzl[idx]
		if idx == replaceIdx {
			val = replaceVal
		}
		cmd := val[:3]
		amt, _ := strconv.Atoi(val[4:])

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
