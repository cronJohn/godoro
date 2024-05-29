package pomomgr

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/cronJohn/godoro/internal/ui"
)

func (pm PomoModel) Init() tea.Cmd {
	return pm.ProgressModel.Init()
}

func (pm PomoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	progressModel, cmd := pm.ProgressModel.Update(msg)
	pm.ProgressModel = progressModel.(ui.ProgressModel)

	switch msg.(type) {
	case tea.QuitMsg:
		pm.cancel()
		return pm, tea.Quit

	case tea.WindowSizeMsg:
		return pm, nil

	case progress.FrameMsg, ui.TickMsg:
		return pm, cmd

	case ui.MaxMsg:
		switch pm.state {
		case WORKING:
			pm.state = BREAKING
			pm.ProgressModel.Total = pm.session.BreakDuration
		case BREAKING:
			pm.state = WORKING
			pm.ProgressModel.Total = pm.session.WorkDuration
		}
		pm.ProgressModel.Current = 0
		resetCmd := pm.ProgressModel.Progress.SetPercent(0)

		return pm, resetCmd
	default:
		return pm, tea.Quit

	}
}
