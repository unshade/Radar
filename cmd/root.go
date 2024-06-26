package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"main/cmd/mode"
	"main/cmd/scan_gateway"
	"main/cmd/scan_ports"
	"main/cmd/version"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "radar",
	Short: "Radar is an easy network scanner and monitoring tool.",
	Long:  "Radar is an easy network scanner and monitoring tool.",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := tea.NewProgram(newModel(), tea.WithAltScreen(), tea.WithMouseCellMotion()).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	Register()
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Register() {
	RootCmd.AddCommand(mode.ExecMode)
	RootCmd.AddCommand(scan_ports.ScanPortsCommand)
	RootCmd.AddCommand(scan_gateway.ScanGatewayCommand)
	RootCmd.AddCommand(version.VersionCmd)
}
