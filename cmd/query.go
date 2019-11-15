package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// QueryCmd to query windows registry
var QueryCmd = &cobra.Command{
	Use:   "QUERY",
	Short: "Queries windows registry",
	Long: `
	Use this command to query windows registry
	
	USAGE: 	reg query -h // prints help
			reg query [key] // queries given key in registry
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Siva Chegondi")
	},
}
