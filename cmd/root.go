package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	tags    string
)

// > app
var rootCmd = &cobra.Command{
	Use:   "godoro",
	Short: "A simple pomodoro CLI written in Go",
	Long:  `Manage your life or something...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome!")

		for _, tag := range strings.Fields(tags) {
			fmt.Printf("Tag: %v\n", tag)
		}
	},
}

// Initialize child commands and flags
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.godoro.yaml)")

	rootCmd.PersistentFlags().
		StringVarP(&tags, "tags", "t", "", "Tags for the session")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".godoro" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".godoro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
