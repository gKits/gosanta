package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	shortExplaination = ""
	longExplaination  = `gosanta is a small cli tool to plan your next secret santa game
with your friends and family. Configure your SMTP server, enter 
all participants names and their email address and get started.`
)

var rootCmd = &cobra.Command{
	Use:   "gosanta",
	Short: shortExplaination,
	Long:  longExplaination,

	Run: func(cmd *cobra.Command, args []string) {
		options := ""
		_ = options
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gosanta.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
