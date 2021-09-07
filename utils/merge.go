package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

const PDFSuffix = ".pdf"

func MergePDF(path, bookName string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Print("failed to read from path", path)
	}
	input := []string{}
	sorter := &NumberSorter{
		files: files,
	}
	err = sorter.sort()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(files); i++ {
		if strings.HasSuffix(files[i].Name(), PDFSuffix) {
			input = append(input, fmt.Sprintf("%s/%s", path, files[i].Name()))
		}
	}
	if !strings.HasSuffix(bookName, ".pdf") {
		bookName = bookName + ".pdf"
	}

	for _, file := range input {
		if err = removeTailPage(file); err != nil {
			panic(err)
		}
	}

	err = mergePDF(input, bookName)
	if err != nil {
		log.Printf("failed to merge PDF files.input: %#v. Error: %v", input, err)
	}
}

// removeTailPage removes the last page of a PDF since it's generally showing something like this:
// For information about citing these materials or our Terms of Use...blah blah...
func removeTailPage(file string) error {
	pageInt, err := pdf.PageCountFile(file)
	if err != nil {
		return err
	}
	return pdf.RemovePagesFile(file, file, []string{fmt.Sprintf("%v", pageInt)}, nil)
}

func mergePDF(files []string, bookName string) error {
	return pdf.MergeCreateFile(files, bookName, nil)
}
