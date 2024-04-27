package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/cronJohn/godoro/util"
)

var (
	FLAG_SECOND, FLAG_MINUTE, FLAG_HOUR int
	FLAG_TAGS                           util.FlagList
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a pomodoro session",
	Long: `Start [godoro start] initiates a new Pomodoro session for a defined duration of time.

Examples:
  # Start a Pomodoro session using default settings
  godoro start

  # Start a 45-minute Pomodoro session with tags 'work' and 'priority'
  godoro start -m 45 -t work,priority
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("Running start...")
		fmt.Printf("Seconds passed in '%v'\n", FLAG_SECOND)
		fmt.Printf("Minutes passed in '%v'\n", FLAG_MINUTE)
		fmt.Printf("Hours passed in '%v'\n", FLAG_HOUR)
		for _, el := range FLAG_TAGS {
			fmt.Printf("Tag: '%v'\n", el)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().
		IntVarP(&FLAG_SECOND, "second", "s", 0, "Specify the duration of the Pomodoro session in seconds. Optional")

	startCmd.Flags().
		IntVarP(&FLAG_MINUTE, "minute", "m", 30, "Specify the duration of the Pomodoro session in minutes. Optional")

	startCmd.Flags().
		IntVarP(&FLAG_HOUR, "hour", "o", 0, "Specify the duration of the Pomodoro session in hours. Optional")

	startCmd.Flags().
		VarP(&FLAG_TAGS, "tags", "t", "Specify the tags of the Pomodoro session. Optional")
}
