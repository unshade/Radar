package scan_gateway

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/cmd"
	"main/internal"
)

func init() {
	cmd.RootCmd.AddCommand(scanGatewayCommand)
}

var scanGatewayCommand = &cobra.Command{
	Use:   "scan-gateways",
	Short: "Scan the gateways of the current network",
	Long:  "Scan the gateways of the current network",
	Run: func(cmd *cobra.Command, args []string) {
		gateways, err := internal.GetGateways()

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, gateway := range gateways {
			fmt.Println(gateway)
			internal.ScanGatewayNetwork(gateway)
		}
	},
}
