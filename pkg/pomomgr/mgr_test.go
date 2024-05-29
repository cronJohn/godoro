package pomomgr

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type TestCase struct {
	workDuration       time.Duration
	breakDuration      time.Duration
	sleepTime          time.Duration
	tags               []string
	expectedTotalWork  time.Duration
	expectedTotalBreak time.Duration
	expectedTags       []string
}

func TestPomoMgrStart(t *testing.T) {
	testCases := []TestCase{
		{
			workDuration:       time.Second * 2,
			breakDuration:      time.Second * 1,
			sleepTime:          time.Second * 7,
			tags:               []string{"test1"},
			expectedTotalWork:  time.Second * 5,
			expectedTotalBreak: time.Second * 2,
			expectedTags:       []string{"test1"},
		},
		{
			workDuration:       time.Second * 1,
			breakDuration:      time.Second * 2,
			sleepTime:          time.Second * 7,
			tags:               []string{"test2"},
			expectedTotalWork:  time.Second * 3,
			expectedTotalBreak: time.Second * 4,
			expectedTags:       []string{"test2"},
		},
		{
			workDuration:       time.Second * 0,
			breakDuration:      time.Second * 0,
			sleepTime:          time.Second * 0,
			tags:               []string{""},
			expectedTotalWork:  time.Second * 0,
			expectedTotalBreak: time.Second * 0,
			expectedTags:       []string{""},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(
			fmt.Sprintf(
				"WorkDuration=%v_BreakDuration=%v_SleepTime=%v_Tags=%v",
				tc.workDuration,
				tc.breakDuration,
				tc.sleepTime,
				tc.tags,
			),
			func(t *testing.T) {
				t.Parallel()
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				pm := NewPomoMgr(tc.workDuration, tc.breakDuration, ctx)

				go pm.Start(tc.tags)

				time.Sleep(tc.sleepTime)

				cancel()

				data := <-pm.Data

				if data.GetWork() != tc.expectedTotalWork {
					t.Errorf(
						"Total work duration: got %v, want %v",
						data.GetWork(),
						tc.expectedTotalWork,
					)
				}

				if data.GetBreak() != tc.expectedTotalBreak {
					t.Errorf(
						"Total break duration: got %v, want %v",
						data.GetBreak(),
						tc.expectedTotalBreak,
					)
				}

				if !reflect.DeepEqual(data.GetTags(), tc.expectedTags) {
					t.Errorf("Tags: got %v, want %v", data.GetTags(), tc.expectedTags)
				}
			},
		)
	}
}
