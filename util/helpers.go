package util

import (
	"io/ioutil"
	"strings"
)

// ReadPuzzleInput - reads puzzle inputs, splits on sep
func ReadPuzzleInput(filenum, sep string) []string {
	pzl, err := ioutil.ReadFile("./inputs/day" + filenum + ".txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(pzl), sep)
}

// ReadRawPuzzleInput - reads puzzle inputs, no splits
func ReadRawPuzzleInput(filenum string) string {
	pzl, err := ioutil.ReadFile("./inputs/day" + filenum + ".txt")
	if err != nil {
		panic(err)
	}
	return string(pzl)
}
