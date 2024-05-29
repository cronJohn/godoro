package ui

import (
	"context"
	"time"

	"github.com/charmbracelet/bubbles/progress"
)

type (
	TickMsg struct{}
	MaxMsg  struct{}
)

type Style struct {
	Padding  int
	MaxWidth int
}

type ProgressModel struct {
	Progress progress.Model
	Current  time.Duration
	Total    time.Duration
	Style    Style
}

func NewProgressModel(total time.Duration, quitFn context.CancelFunc) ProgressModel {
	return ProgressModel{
		Progress: progress.New(progress.WithDefaultGradient()),
		Current:  0,
		Total:    total,
		Style:    Style{Padding: 3, MaxWidth: 50},
	}
}

func (p ProgressModel) WithStyle(style Style) ProgressModel {
	p.Style = style
	return p
}
