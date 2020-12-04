package util

import (
	"io/ioutil"
	"strings"
)

// ReadPuzzleInput - reads puzzle inputs, obv.
func ReadPuzzleInput(filenum, sep string) []string {
	pzl, err := ioutil.ReadFile("./inputs/day" + filenum + ".txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(pzl), sep)
}
