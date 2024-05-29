package pomomgr

import (
	"context"
	"testing"
	"time"
)

func TestPomoMgrStart(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	pm := NewPomoMgr(time.Second*2, time.Second*1, ctx)

	go func() {
		pm.Start([]string{"test"})
	}()

	time.Sleep(time.Second * 7)

	cancel()

	if pm.Data.GetWork() != time.Second*5 {
		t.Errorf("data.totalWork = %v, want %v", pm.Data.totalWork, time.Second*5)
	}

	if pm.Data.GetBreak() != time.Second*2 {
		t.Errorf("data.totalBreak = %v, want %v", pm.Data.totalBreak, time.Second*2)
	}
}
