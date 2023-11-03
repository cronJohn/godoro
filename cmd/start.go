package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/cronJohn/godoro/types/timeutil"
)

var secs, mins, hours int

const (
	TimerMode int = iota
	StopwatchMode
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a pomodoro session",
	Long:  `Use this command to just start working...`,
	Run: func(cmd *cobra.Command, args []string) {
		var mode int

		if cmd.Flags().Changed("seconds") || cmd.Flags().Changed("minutes") ||
			cmd.Flags().Changed("hours") {
			mode = TimerMode
		} else {
			mode = StopwatchMode
		}

		switch mode {
		case TimerMode:
			log.Info().
				Int("seconds", secs).
				Int("minutes", mins).
				Int("hours", hours).
				Msg("Timer mode")

			timer := timeutil.NewTimer(
				time.Duration(secs)*time.Second,
				time.Duration(mins)*time.Minute,
				time.Duration(hours)*time.Hour,
			)
			timer.Start()
		case StopwatchMode:
			log.Info().Msg("Stopwatch mode")
			sw := timeutil.NewStopwatch()
			sw.Start()
			interruptChan := make(chan os.Signal, 1)
			signal.Notify(interruptChan, os.Interrupt)

			<-interruptChan // Block until we receive a signal on the channel

			// Print the elapsed time with 2 decimal places
			fmt.Printf("Elapsed time: %.2f seconds\n", sw.Stop().Seconds())

		default:
			log.Fatal().Msg("Unknown time mode...")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntVarP(&secs, "seconds", "s", 0, "Number of seconds to run the session")
	startCmd.Flags().IntVarP(&mins, "minutes", "m", 0, "Number of minutes to run the session")
	startCmd.Flags().IntVarP(&hours, "hours", "o", 0, "Number of hours to run the session")
}
