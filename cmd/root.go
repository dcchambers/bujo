package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bujo",
	Short: "Bujo is a command line bullet journal.",
	Long: `A bullet journal on your command line.
				  Learn more about bullet journaling: https://bulletjournal.com
				  Learm more about bujo: https://github.com/dcchambers/bujo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello bujo!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
