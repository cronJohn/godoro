package cmd

import (
	"context"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/cronJohn/godoro/pkg/pomomgr"
	"github.com/cronJohn/godoro/util"
)

var (
	FLAG_WORK_DURATION, FLAG_BREAK_DURATION time.Duration
	FLAG_TAGS                               util.TagList
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a pomodoro session",
	Long: `Start [godoro start] initiates a new Pomodoro session for a defined duration of time.

Examples:
  # Start a Pomodoro session using default settings
  godoro start

  # Start a 25-minute work and 5-minute break Pomodoro session with tags 'work' and 'priority'
  godoro start -w 25m -b 5m -t work,priority
`,
	Run: handleCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().
		DurationVarP(&FLAG_WORK_DURATION, "work", "w", time.Minute*25, "Specify the work duration of the Pomodoro session. Optional")

	startCmd.Flags().
		DurationVarP(&FLAG_BREAK_DURATION, "break", "b", time.Minute*5, "Specify the break duration of the Pomodoro session. Optional")

	startCmd.Flags().
		VarP(&FLAG_TAGS, "tags", "t", "Specify the tags of the Pomodoro session. Optional")
}

func handleCmd(cmd *cobra.Command, args []string) {
	log.Debug().Msg("Running start...")

	pm := pomomgr.NewPomoMgr(FLAG_WORK_DURATION, FLAG_BREAK_DURATION, context.Background())

	go func() {
		pm.Start(FLAG_TAGS)
	}()

	m := pomomgr.NewProgressModel(pomomgr.PomoSession{
		WorkDuration:  FLAG_WORK_DURATION,
		BreakDuration: FLAG_BREAK_DURATION,
		State:         pomomgr.WORKING,
	})

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Error().Msg("Something went wrong...")
		os.Exit(1)
	}
}
