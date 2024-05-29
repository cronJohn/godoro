package pomomgr

import (
	"strings"
)

func (m PomoModel) View() string {
	pad := strings.Repeat(" ", m.ProgressModel.Style.Padding)
	var message string
	switch m.state {
	case WORKING:
		message = "Still working.."
	case BREAKING:
		message = "Still on break.."
	}

	return "\n" +
		pad + message + "\n\n" +
		pad + m.ProgressModel.View() + "\n\n"
}
