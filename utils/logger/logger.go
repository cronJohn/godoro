package logger

import (
	"log"
)

type LoggerAggregate struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

func (l *LoggerAggregate) Info(v ...interface{}) {
	l.InfoLogger.Println(v...)
}

func (l *LoggerAggregate) Warning(v ...interface{}) {
	l.WarningLogger.Println(v...)
}

func (l *LoggerAggregate) Error(v ...interface{}) {
	l.ErrorLogger.Println(v...)
}
