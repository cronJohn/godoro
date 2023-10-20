package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"

	"github.com/cronJohn/godoro/utils/logger"
)

var flags int = log.Ldate | log.Ltime | log.Lshortfile

func main() {
	fmt.Println("Welcome!")

	lg := logger.LoggerAggregate{
		InfoLogger:    log.New(os.Stdout, color.CyanString("[INFO] "), flags),
		WarningLogger: log.New(os.Stdout, color.YellowString("[WARNING] "), flags),
		ErrorLogger:   log.New(os.Stderr, color.RedString("[ERROR] "), flags),
	}

	lg.Info("This is an info message")
	lg.Warning("This is a warning message")
	lg.Error("This is an error message")
}
