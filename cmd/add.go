package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AddCmd to query windows registry
var AddCmd = &cobra.Command{
	Use:   "ADD",
	Short: "Add new keys to Windows Registry",
	Long: `
Use this command to add new key windows registry
	
USAGE:
	reg ADD -h // prints help
	reg ADD [key] [value] // add given key value to registry
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Siva Chegondi")
	},
}
