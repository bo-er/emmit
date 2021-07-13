package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bo-er/emmit/links"
	"github.com/bo-er/emmit/utils"
	"github.com/spf13/cobra"
)

const EmptyString = ""

var pdfCommand = &cobra.Command{
	Short: "generate your OCW course book(if there is any course materials)",
	Use:   "pdf",
	Long: `pdf command needs three input
			1. first is the file where you store all the links downloaded by the crawl command
			2. second is the course number. e.g, 6-034
			3. third is the year (if the course is given during multiple years, leave it empty if you want course materials of the latest year)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Panic(`Dear Foolio,I need at least 2 parameters, the first one is file where you store all the links downloaded by the crawl command,
			the second one is the course number`)
		}
		var pivot int
		var courseNumber = args[1]
		var yearUncertain bool = len(args) < 3
		var year string
		if len(args) >= 3 {
			year = args[2]
		}
		courses := make(map[string][]string)
		lines, err := utils.Readline(args[0])
		if err != nil {
			log.Panic(err)
		}

		// doing all the filter jobs
		for _, line := range lines {
			if strings.Contains(line, courseNumber) {
				y := links.GetLinkYear(line)
				if isYearMatch(y, year) && isCourseLectures(line) {
					courses[y] = append(courses[y], line)
				}
			}
		}

		wanted := getWantedLinks(yearUncertain, pivot, courses)
		for _, p := range wanted {
			name := links.GetPDFName(p)
			if name == EmptyString {
				continue
			}

			path := fmt.Sprintf("./%s/%s", courseNumber, name)

			utils.Download(path, p)
		}
	},
}

func getWantedLinks(yearUncertain bool, pivot int, courses map[string][]string) []string {
	if yearUncertain && len(courses) == 1 {
		for _, v := range courses {
			return v
		}
	} else if yearUncertain {
		for k, _ := range courses {
			i, err := strconv.Atoi(k)
			if err != nil {
				log.Panic(err)
			}
			if pivot == 0 || i > pivot {
				pivot = i
			}
		}
	}
	if v, ok := courses[strconv.Itoa(pivot)]; ok {
		return v
	}
	return nil
}

func isYearMatch(y, year string) bool {
	if year == "" {
		return true
	}
	return year == y //year is not empty
}

func isCourseLectures(link string) bool {
	return strings.Contains(link, "lecture-slides") && strings.Contains(link, ".pdf")
}
