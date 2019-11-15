package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
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
		} else {
			k, err := registry.OpenKey(registry.LOCAL_MACHINE, args[0], registry.QUERY_VALUE)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Printf("Value of key %s is %+v\n", args[0], k)
			}
		}
	},
}
