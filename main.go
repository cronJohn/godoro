package main

import (
	"log"
	"os"

	"github.com/fatih/color"

	"github.com/cronJohn/godoro/cmd"
	"github.com/cronJohn/godoro/utils/logger"
)

var flags int = log.Ldate | log.Ltime | log.Lshortfile

func init() {
	Lg := logger.LoggerAggregate{
		InfoLogger:    log.New(os.Stdout, color.CyanString("[INFO] "), flags),
		WarningLogger: log.New(os.Stdout, color.YellowString("[WARNING] "), flags),
		ErrorLogger:   log.New(os.Stderr, color.RedString("[ERROR] "), flags),
	}

	Lg.Info("Logger created...")
}

func main() {
	cmd.Execute()
}
