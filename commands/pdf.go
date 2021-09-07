package commands

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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
			3. third is the year (if the course is given during multiple years, leave it empty if you want course materials of the latest year)
			4. fourth is the alternative keyword in case the pdf notes have different name. e.g., 
			"6-042j-mathematics-for-computer-science-fall-2010/recitations/MIT6_042JF10_rec03_sol.pdf" -> in this case name is rec01,rec02,etc.
			`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Panic(`Dear Foolio,I need at least 2 parameters, the first one is file where you store all the links downloaded by the crawl command,
			the second one is the course number`)
		}
		var courseNumber = args[1]
		var year string
		var pattern string
		var reg *regexp.Regexp
		if len(args) >= 3 {
			year = args[2]
		}
		if len(args) >= 4 {
			pattern = args[3]
			reg = regexp.MustCompile(pattern)
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

				if pattern != "" { // then using custom patterns
					if isYearMatch(y, year) && filterByRegex(line, reg) {
						courses[y] = append(courses[y], line)
					}
				} else {
					if isYearMatch(y, year) && isCourseLectures(line) {
						courses[y] = append(courses[y], line)
					}
				}

			}
		}

		wanted := getWantedLinks(year, courses)
		for _, p := range wanted {
			name := links.GetPDFName(p)
			if name == EmptyString {
				continue
			}
			err := os.MkdirAll(fmt.Sprintf("./%s", courseNumber), 0755)
			if err != nil {
				panic(fmt.Errorf("failed to call os.MkdirAll. Error: %w", err))
			}
			path := fmt.Sprintf("./%s/%s", courseNumber, name)
			utils.Download(path, p)
		}

		utils.MergePDF(fmt.Sprintf("./%s", courseNumber), fmt.Sprintf("MIT%s.pdf", courseNumber))
	},
}

func getCorseMapKeys(m map[string][]string) []string {
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func getWantedLinks(year string, courses map[string][]string) []string {
	if year != "" {
		if links, exists := courses[year]; exists {
			return links
		}
	} else {
		keys := getCorseMapKeys(courses)
		sort.Strings(keys)
		return courses[keys[len(keys)-1]]
	}
	return nil
}

func isYearMatch(y, year string) bool {
	if year == "" {
		return true
	}
	return year == y //year is not empty
}

// isCourseLectures filters lecture link
func isCourseLectures(link string) bool {
	return (strings.Contains(link, "lecture-slides") ||
		strings.Contains(link, "lecture-notes")) ||
		strings.Contains(link, "lecture-videos") ||
		strings.Contains(link, ".pdf")
}

// filterByRegex filters a link whose suffix matches the regex expression
func filterByRegex(link string, keyword *regexp.Regexp) bool {
	// filter non pdf link
	if !strings.HasSuffix(link, ".pdf") {
		return false
	}
	if keyword != nil {
		match := keyword.FindString(link)
		if match == "" {
			return false
		}
		return strings.HasSuffix(link, match)
	}

	return true
}
