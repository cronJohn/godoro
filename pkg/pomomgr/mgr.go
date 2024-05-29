package pomomgr

import (
	"context"
	"time"
)

// PomoSession represents data related to how long the pomodoro work
// and break session should be
type PomoSession struct {
	WorkDuration  time.Duration
	BreakDuration time.Duration
}

// PomoData represents data related to the pomodoro session
// This includes the total work and break time
// and the tag associated with the session
type PomoData struct {
	totalWork  time.Duration
	totalBreak time.Duration
	tag        []string
}

// Returns a rounded version of totalWork in PomoData
func (p PomoData) GetWork() time.Duration {
	return p.totalWork.Round(time.Second)
}

// Returns a rounded version of totalBreak in PomoData
func (p PomoData) GetBreak() time.Duration {
	return p.totalBreak.Round(time.Second)
}

// PomoMgr is a Pomodoro session manager
// It starts a Pomodoro session
// and manages PomoData associated with the session
//
// It can be stopped using its context quit field
type PomoMgr struct {
	session PomoSession
	Data    PomoData
	quit    context.Context
}

func NewPomoMgr(workDuration, breakDuration time.Duration, ctx context.Context) *PomoMgr {
	return &PomoMgr{
		session: PomoSession{
			WorkDuration:  workDuration,
			BreakDuration: breakDuration,
		},
		Data: PomoData{},
		quit: ctx,
	}
}

func (pm *PomoMgr) Start(tag []string) {
	var timeBuf time.Time
	var data struct {
		totalWork  time.Duration
		totalBreak time.Duration
	}

	for {
		timeBuf = time.Now()
		select { // work
		case <-pm.quit.Done():
			data.totalWork += time.Since(timeBuf)
			pm.Data = PomoData{
				totalWork:  data.totalWork,
				totalBreak: data.totalBreak,
				tag:        tag,
			}
			return
		case <-time.After(pm.session.WorkDuration):
			data.totalWork += pm.session.WorkDuration

			timeBuf = time.Now()
			select { // break
			case <-pm.quit.Done():
				data.totalBreak += time.Since(timeBuf)
				pm.Data = PomoData{
					totalWork:  data.totalWork,
					totalBreak: data.totalBreak,
					tag:        tag,
				}
				return
			case <-time.After(pm.session.BreakDuration):
				data.totalBreak += pm.session.BreakDuration
				break
			}
		}
	}
}
