/*
Copyright © 2024 donclaveau3

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ByteCount bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "This is a custom version of wc for a coding challenge",
	Long: `ccwc is a CLI app that imitates the behavior of wc.
This application is a solution to a coding challenge found here:
https://codingchallenges.fyi/challenges/challenge-wc`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]
		if ByteCount {
			message = fmt.Sprintf("bytecount %s", args[0])
		}
		fmt.Println(message)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&ByteCount, "bytecount", "c", false, "include byte count for file contents")
}