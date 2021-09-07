package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "This is a test command",
	Long:  `this is a test command for doing anything`,
	Run: func(cmd *cobra.Command, args []string) {
		// links.StartCrawlingFilteringFunc([]string{"http://blog.yufeng.info/page/1"}, "./blog4.info", 500*time.Millisecond, FilterBlog)
		// commentsMap := make(map[string][]string)

		// files, err := ioutil.ReadDir("./blog")
		// if err != nil {
		// 	panic(err)
		// }
		// var reg = regexp.MustCompile(`http:\/\/blog.yufeng.info\/archives\/([0-9]+)#comment-([0-9]+)`)
		// for _, f := range files {
		// 	content, err := ioutil.ReadFile("./blog/"+f.Name())
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	results := reg.FindAllStringSubmatch(string(content), -1)
		// }
	},
}

func FilterBlog() func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, "http://blog.yufeng.info")
	}
}
