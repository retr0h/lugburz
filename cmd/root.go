package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "lugctl",
	Short: "A CLI to control the Lugbúrz service.",
	Long: `
 _             _   _
| |_ _ ___ ___| |_| |
| | | | . |  _|  _| |
|_|___|_  |___|_| |_|
      |___|

A command line tool which controls the Lugbúrz service.

https://github.com/retr0h/lugburz
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
