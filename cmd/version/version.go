package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Radar",
	Long:  `Print the version number of Radar`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Radar v0.1")
	},
}
