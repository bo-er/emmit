package commands

import (
	"fmt"
	"sync"

	"github.com/bo-er/emmit/utils"
	"github.com/spf13/cobra"
)

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "every good commander should have his test command, so do I.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.MergePDF("./6-006", "MIT6-006.pdf")
	},
}
var once sync.Once

func test() {
	once.Do(func() {
		fmt.Println("hhh")
	})
}
