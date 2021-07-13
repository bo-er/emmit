package commands

import (
	"github.com/bo-er/emmit/utils"
	"github.com/spf13/cobra"
)

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "every good commander should have his test command, so do I.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.MergePDF("./6_172", "6_172")
	},
}
