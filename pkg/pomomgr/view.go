package pomomgr

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	helpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	stateStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF00FF")).Render
	progressStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500")).Render
	borderStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)
)

func (m ProgressModel) View() string {
	var stateText string

	switch m.Session.State {
	case WORKING:
		stateText = stateStyle("Working")
	case BREAKING:
		stateText = stateStyle("On Break")
	}

	content := lipgloss.JoinVertical(lipgloss.Left,
		titleStyle("Pomodoro Timer"),
		"",
		progressStyle(m.Progress.View()),
		"",
		stateText,
		"",
		helpStyle("Press any key to quit"),
	)

	return borderStyle.Render(content)
}
