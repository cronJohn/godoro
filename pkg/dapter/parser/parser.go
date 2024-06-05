package parser

import "github.com/cronJohn/godoro/pkg/dapter/proto"

type Parser interface{}

type PomoParser interface {
	GetSessions() (*proto.PomoSessions, error)
}
