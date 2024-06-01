package pomomgr

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	WORKING = iota
	BREAKING
)

// PomoSession represents data related to how long the pomodoro work
// and break session should be
type PomoSession struct {
	WorkDuration  time.Duration
	BreakDuration time.Duration
	State         int
}

// PomoData represents data related to the pomodoro session
// This includes the total work and break time
// and the tag associated with the session
type PomoData struct {
	totalWork  time.Duration
	totalBreak time.Duration
	tags       []string
}

// Returns a rounded version of totalWork in PomoData
func (p PomoData) GetWork() time.Duration {
	return p.totalWork.Round(time.Second)
}

// Returns a rounded version of totalBreak in PomoData
func (p PomoData) GetBreak() time.Duration {
	return p.totalBreak.Round(time.Second)
}

// Returns the tags in PomoData
func (p PomoData) GetTags() []string {
	return p.tags
}

// PomoMgr is a Pomodoro session manager
// It starts a Pomodoro session
// and manages PomoData associated with the session
//
// It can be stopped using its context quit field
type PomoMgr struct {
	session PomoSession
	Data    chan PomoData
	quit    context.Context
}

func NewPomoMgr(workDuration, breakDuration time.Duration, ctx context.Context) *PomoMgr {
	return &PomoMgr{
		session: PomoSession{
			WorkDuration:  workDuration,
			BreakDuration: breakDuration,
		},
		Data: make(chan PomoData),
		quit: ctx,
	}
}

func (pm *PomoMgr) Start(tags []string) {
	var timeBuf time.Time
	var data PomoData

	data.tags = tags

	for {
		timeBuf = time.Now()
		log.Debug().Msg("Starting work")
		select { // work
		case <-pm.quit.Done():
			data.totalWork += time.Since(timeBuf)
			pm.Data <- data
			return
		case <-time.After(pm.session.WorkDuration):
			data.totalWork += pm.session.WorkDuration

			timeBuf = time.Now()
			log.Debug().Msg("Starting break")
			select { // break
			case <-pm.quit.Done():
				data.totalBreak += time.Since(timeBuf)
				pm.Data <- data
				return
			case <-time.After(pm.session.BreakDuration):
				data.totalBreak += pm.session.BreakDuration
				break
			}
		}
	}
}
