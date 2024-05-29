package pomomgr

import (
	"context"

	"github.com/cronJohn/godoro/internal/ui"
)

const (
	WORKING = iota
	BREAKING
)

type PomoModel struct {
	ProgressModel ui.ProgressModel
	session       PomoSession
	state         byte
	cancel        context.CancelFunc
}

func NewPomoModel(session PomoSession, quitFn context.CancelFunc) PomoModel {
	return PomoModel{
		ProgressModel: ui.NewProgressModel(session.WorkDuration, quitFn),
		session:       session,
		state:         WORKING,
		cancel:        quitFn,
	}
}
