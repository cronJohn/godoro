package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "godoro",
	Short: "A simple CLI pomodoro tool written in Go",
	Long: `                   .___                   
   ____   ____   __| _/___________  ____  
  / ___\ /  _ \ / __ |/  _ \_  __ \/  _ \ 
 / /_/  >  <_> ) /_/ (  <_> )  | \(  <_> )
 \___  / \____/\____ |\____/|__|   \____/ 
/_____/             \/                    
A CLI tool that let's you create and manage
multiple pomodoro sessions. Additionally,
you can view stats about your different
sessions as well as set tags for each session.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error().Msg("Something went wrong with cobra...")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.godoro.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Error().Msg("Unable to get user's home directory")
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".godoro")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Warn().Msgf("Using config file: %s", viper.ConfigFileUsed())
	}
}
