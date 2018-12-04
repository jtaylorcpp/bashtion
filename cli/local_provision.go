package cli

import (
	"fmt"

	"github.com/jtaylorcpp/bashtion/provisioning"
	"github.com/spf13/cobra"
)

func init() {
	LocalProvisionCommand.Flags().StringVarP(&FileLocation, "file", "f", "", "File to provision from.")
}

var FileLocation string

var LocalProvisionCommand = &cobra.Command{
	Use:   "local",
	Short: "Provision using a local script.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Local provision from file: %s\n", FileLocation)
		provisioning.ProvisionFromFile(FileLocation)
	},
}
