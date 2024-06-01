package pomomgr

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 80
)

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - padding*2 - 4
		if m.Progress.Width > maxWidth {
			m.Progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.Progress.Percent() >= 1.0 {
			switch m.Session.State {
			case WORKING:
				return m, changeStateCmd(BREAKING)
			case BREAKING:
				return m, changeStateCmd(WORKING)
			}
		}

		m.runningTotal += time.Second

		cmd := m.Progress.SetPercent(float64(m.runningTotal) / float64(m.stateToTotalMap[m.Session.State]))
		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	case changeStateMsg:
		m.Session.State = msg.state
		m.Reset()
		return m, tickCmd()
	}

	return m, nil
}
