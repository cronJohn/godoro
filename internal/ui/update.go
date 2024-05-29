package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m ProgressModel) Init() tea.Cmd {
	return TickCmd()
}

func TickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}

func MaxCmd() tea.Cmd {
	return func() tea.Msg {
		return MaxMsg{}
	}
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.QuitMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - m.Style.Padding*2 - 4
		if m.Progress.Width > m.Style.MaxWidth {
			m.Progress.Width = m.Style.MaxWidth
		}
		return m, nil

	case TickMsg:
		if m.Progress.Percent() >= 1.0 {
			return m, MaxCmd()
		}
		m.Current += 1 * time.Second
		cmd := m.Progress.SetPercent(float64(m.Current) / float64(m.Total))
		return m, cmd

	default:
		return m, nil
	}
}
