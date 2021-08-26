package utils

import (
	"fmt"
	"io/fs"
	"sort"
	"strconv"
	"strings"
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

const Lecprefix = "lec"

func (ns *NumberSorter) sort() error {
	for i := 0; i < len(ns.files); i++ {
		number := strings.TrimRight(strings.TrimPrefix(ns.files[i].Name(), Lecprefix), PDFSuffix)
		_, err := strconv.Atoi(number)
		if err != nil {
			return err
		}
	}
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
	numberI := strings.TrimRight(strings.TrimPrefix(n.files[i].Name(), Lecprefix), PDFSuffix)
	indexI, _ := strconv.Atoi(numberI)
	numberJ := strings.TrimRight(strings.TrimPrefix(n.files[j].Name(), Lecprefix), PDFSuffix)
	indexJ, _ := strconv.Atoi(numberJ)

	return indexI < indexJ
}
