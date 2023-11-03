package timeutil

import "time"

type Stopwatch struct {
	startTime time.Time
	isRunning bool
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		startTime: time.Time{},
		isRunning: false,
	}
}

func (s *Stopwatch) Start() {
	if !s.isRunning {
		s.startTime = time.Now()
		s.isRunning = true
	}
}

func (s *Stopwatch) Stop() time.Duration {
	if s.isRunning {
		elapsed := time.Since(s.startTime)
		s.isRunning = false
		return elapsed
	}
	return 0
}
