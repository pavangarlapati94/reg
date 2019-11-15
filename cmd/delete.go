package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// DeleteCmd to query windows registry
var DeleteCmd = &cobra.Command{
	Use:   "DELETE",
	Short: "Deletes windows registry",
	Long: `
Use this command to query windows registry
	
USAGE:
	reg DELETE -h // prints help
	reg DELETE [key] // deletes given key in registry
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
		}
		fmt.Println(args);
	},
}
