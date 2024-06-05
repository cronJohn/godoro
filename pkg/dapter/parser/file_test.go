package parser

import (
	"reflect"
	"testing"
	"time"

	"github.com/cronJohn/godoro/pkg/dapter/proto"
)

type TestTable struct {
	Id            int32
	WorkDuration  time.Duration
	BreakDuration time.Duration
	Tags          []string
}

var testCases = []TestTable{
	{
		Id:            1,
		WorkDuration:  time.Second * 1,
		BreakDuration: time.Second * 2,
		Tags:          []string{"foo"},
	},
	{
		Id:            2,
		WorkDuration:  time.Second * 3,
		BreakDuration: time.Second * 4,
		Tags:          []string{"bar", "baz"},
	},
}

func compareSessions(t *testing.T, sessions *proto.PomoSessions) {
	for i, val := range sessions.GetSessions() {
		if val.GetId() != testCases[i].Id {
			t.Errorf("expected id %v, got %v", testCases[i].Id, val.GetId())
		}

		if val.GetWorkDuration().AsDuration() != testCases[i].WorkDuration {
			t.Errorf(
				"expected work duration %v, got %v",
				testCases[i].WorkDuration,
				val.GetWorkDuration().AsDuration(),
			)
		}

		if val.GetBreakDuration().AsDuration() != testCases[i].BreakDuration {
			t.Errorf(
				"expected break duration %v, got %v",
				testCases[i].BreakDuration,
				val.GetBreakDuration().AsDuration(),
			)
		}

		if !reflect.DeepEqual(val.GetTags(), testCases[i].Tags) {
			t.Errorf(
				"expected tags %v, got %v",
				testCases[i].Tags,
				val.GetTags(),
			)
		}
	}
}

func TestReadCSV(t *testing.T) {
	parser := NewFileParser(CSV, "test.csv")

	sessions, err := parser.GetSessions()
	if err != nil {
		t.Error(err)
	}

	compareSessions(t, sessions)
}

func TestReadJSON(t *testing.T) {
	parser := NewFileParser(JSON, "test.json")

	sessions, err := parser.GetSessions()
	if err != nil {
		t.Error(err)
	}

	compareSessions(t, sessions)
}
