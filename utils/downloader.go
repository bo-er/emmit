package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Download(path, url string) {
	out, err := os.Create(path)
	if err != nil {
		if _, isPathError := err.(*os.PathError); isPathError {
			index := strings.LastIndex(path, "/")

			err := os.Mkdir(path[:index], 0755)
			if err != nil {
				log.Print("failed to create dir.err:", err)
			}
			out, err = os.Create(path)
			if err != nil {
				panic(err)
			}
		}
	} else {
		log.Print(err)
	}

	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Print(err)
	}
}
