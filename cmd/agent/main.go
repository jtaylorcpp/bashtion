package main

import (
	"fmt"

	"github.com/jtaylorcpp/bashtion/cli"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cli.LocalProvisionCommand)
}

var rootCmd = &cobra.Command{
	Use:   "bashtion-agent",
	Short: "Bashtion Agent is the thing that installs things",
	Long: `Bashtion Agent is a binary that will run and
	return output from BASH scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error from agent: %v\n", err)
	}
}
