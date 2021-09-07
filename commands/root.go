package commands

import (
	"fmt"
	"os"

	"github.com/bo-er/emmit/links"
	"github.com/spf13/cobra"
)

var ParentUrl string
var School string //every OCW link contains a shcool name like `electrical-engineering-and-computer-science`, if you don't
var File string

func init() {
	RootCmd.PersistentFlags().StringVarP(&ParentUrl, "parenturl", "p", "https://ocw.mit.com", "this is the parent url of emmit")
	RootCmd.PersistentFlags().StringVarP(&File, "file", "f", "links.info", "this is the file that all the links are stored at")
	RootCmd.PersistentFlags().StringVarP(&School, "school", "s", "electrical-engineering-and-computer-science",
		"this is the school you're interested in, default value is EE school")
	RootCmd.AddCommand(pdfCommand)
	RootCmd.AddCommand(testCmd)
}

var RootCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl any website that are from the main site https://ocw.mit.edu",
	Long:  `Crawl any website that are from the main site https://ocw.mit.edu, if you don't pass in the "s" argument, EE school is picked`,
	Run: func(cmd *cobra.Command, args []string) {
		links.StartCrawlingFilteringString([]string{ParentUrl}, File, School)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
