package dapter

import (
	"github.com/cronJohn/godoro/pkg/dapter/parser"
	"github.com/cronJohn/godoro/pkg/dapter/proto"
)

type PomoDapter struct {
	Parser parser.PomoParser
}

func (d *PomoDapter) GetSessions() (*proto.PomoSessions, error) {
	data, err := d.Parser.GetSessions()
	return data, err
}

// func (d PomoDapter) WriteSession(data *proto.PomoSession) error {
// 	err := d.Parser.WriteSession(data)
// 	return err
// }
