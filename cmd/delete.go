package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// DeleteQuery to query windows registry
var DeleteQuery = &cobra.Command{
	Use:   "DELETE",
	Short: "Deletes windows registry",
	Long: `
Use this command to query windows registry
	
USAGE:
	reg DELETE -h // prints help
	reg DELETE [key] // deletes given key in registry
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Siva Chegondi")
	},
}
