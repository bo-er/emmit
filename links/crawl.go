package links

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	OCW   = "ocw.mit.edu"
	OCWEE = "https://ocw.mit.edu/courses/electrical-engineering-and-computer-science"
)

type FilterFunc func() func(string) bool

// FoundCourseSlides is used to filter out links that are not lecture slides
func FoundCourseSlides() func(string) bool {
	return func(url string) bool {
		return strings.HasSuffix(url, "lecture-slides/index.htm") ||
			strings.HasSuffix(url, "lecture-slides/index.html")
	}
}

// IsMITWebsite makes sure that the crawler does not look beyond OCW
func IsMITWebsite() func(string) bool {
	return func(url string) bool {
		return strings.Contains(url, OCW)
	}
}

// IsEESchool is used to filter out links that are not from the EE school.
func IsEESchool() func(string) bool {
	return func(url string) bool {
		return strings.HasPrefix(url, OCWEE)
	}
}

// tokens is used as a concurrent requests limiter
var tokens = make(chan struct{}, 20)

func crawlWithFilterFunc(w io.Writer, url string, filterFuncs ...FilterFunc) []string {
	returnList := []string{}
	fmt.Fprintln(w, url)
	tokens <- struct{}{} //using token to limit concurrent requests, in this case it's 20
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}

	filter := func(link string, filterFuncs ...FilterFunc) {
		for _, ff := range filterFuncs {
			if !ff()(link) {
				return
			}
		}
		returnList = append(returnList, link)
	}

	for _, link := range list {
		filter(link, filterFuncs...)
	}
	return returnList
}

func crawlWithFilterString(w io.Writer, url, filterString string) []string {
	returnList := []string{}
	fmt.Fprintln(w, url)
	tokens <- struct{}{} //using token to limit concurrent requests, in this case it's 20
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}

	filter := func(link, filterString string) {
		if !strings.Contains(link, filterString) {
			return
		}
		returnList = append(returnList, link)
	}

	for _, link := range list {
		filter(link, filterString)
	}
	return returnList
}

func StartCrawlingFilteringFunc(parentUrls []string, filePath string, filterFuncs ...FilterFunc) {
	var file io.Writer
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			log.Print(err)
		}
	}
	worklist := make(chan []string)
	var n int = 10 //number of pending sends to worklist
	n++
	go func() {
		worklist <- parentUrls
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawlWithFilterFunc(file, link, filterFuncs...)
				}(link)
			}
		}
	}

}

func StartCrawlingFilteringString(parentUrls []string, filePath, filterString string) {
	var file io.Writer
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			log.Print(err)
		}
	}
	worklist := make(chan []string)
	var n int = 10 //number of pending sends to worklist
	n++
	go func() {
		worklist <- parentUrls
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawlWithFilterString(file, link, filterString)
				}(link)
			}
		}
	}

}
