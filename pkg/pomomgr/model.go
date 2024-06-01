package pomomgr

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	tickMsg time.Time
)

type changeStateMsg struct {
	state int
}

type ProgressModel struct {
	Session         PomoSession
	Progress        progress.Model
	stateToTotalMap map[int]time.Duration
	runningTotal    time.Duration
}

func NewProgressModel(session PomoSession) ProgressModel {
	return ProgressModel{
		Session:  session,
		Progress: progress.New(progress.WithDefaultGradient()),
		stateToTotalMap: map[int]time.Duration{
			WORKING:  session.WorkDuration,
			BREAKING: session.BreakDuration,
		},
		runningTotal: 0,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return tickCmd()
}

func (m *ProgressModel) Reset() {
	m.runningTotal = 0
	m.Progress.SetPercent(0.0)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func changeStateCmd(state int) tea.Cmd {
	return func() tea.Msg {
		return changeStateMsg{
			state: state,
		}
	}
}
