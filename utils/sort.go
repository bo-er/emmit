package utils

import (
	"io/fs"
	"regexp"
	"sort"
	"strconv"
)

// Sort dictates in what order files are merged
type Sort interface {
	sort()
}

// NumberSorter sorts by pdf file's suffix that is a number.
// lec1, lec2, lec3, etc.
type NumberSorter struct {
	files []fs.DirEntry
}

func getNumber(filename string) int {
	results := numReg.FindAll([]byte(filename), -1)
	number := results[len(results)-1]
	if number[0] == byte(0) {
		number = number[1:]
	}
	n, err := strconv.Atoi(string(number))
	if err != nil {
		panic(err)
	}
	return n
}

var numReg = regexp.MustCompile("[0-9]+")

func (ns *NumberSorter) sort() error {
	sort.Sort(ns)
	return nil
}

func (n NumberSorter) Len() int {
	return len(n.files)
}
func (n NumberSorter) Swap(i, j int) {
	n.files[i], n.files[j] = n.files[j], n.files[i]
}
func (n NumberSorter) Less(i, j int) bool {
	return getNumber(n.files[i].Name()) < getNumber(n.files[j].Name())
}
