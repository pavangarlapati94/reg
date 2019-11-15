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
	
USAGE:
	reg QUERY -h // prints help
	reg QUERY [key] // queries given key in registry
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			cmd.Help()
		}
		fmt.Println("Siva Chegondi")
	},
}
