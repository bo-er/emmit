package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	// To get your free API key for metered license, sign up on: https://cloud.unidoc.io
	// Make sure to be using UniPDF v3.19.1 or newer for Metered API key support.
	err := license.SetMeteredKey(`99237715ac17115d5265cd2738e050e07d74a9c74375360d2f4d8b8467ce1797`)
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		panic(err)
	}
}

func MergePDF(path, bookName string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Print(err)
	}
	input := []string{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(),".pdf"){
			input = append(input, fmt.Sprintf("%s/%s", path, f.Name()))
		}
	}
	if !strings.HasSuffix(bookName, ".pdf") {
		bookName = bookName + ".pdf"
	}
	err = mergePdf(input, bookName)
	if err != nil {
		log.Print(err)
	}
}

func mergePdf(inputPaths []string, outputPath string) error {
	pdfWriter := model.NewPdfWriter()

	for _, inputPath := range inputPaths {
		pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
		if err != nil {
			return err
		}
		defer f.Close()

		numPages, err := pdfReader.GetNumPages()
		if err != nil {
			return err
		}

		for i := 0; i < numPages; i++ {
			pageNum := i + 1

			page, err := pdfReader.GetPage(pageNum)
			if err != nil {
				return err
			}

			err = pdfWriter.AddPage(page)
			if err != nil {
				return err
			}
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}
