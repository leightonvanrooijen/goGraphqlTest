// Package logger provides Info, Warn and Error loggers
// all logging to logs.txt file. Extends the Golangs log package
package logger

import (
	"log"
	"os"
)

type logger struct {
	logger log.Logger
}

var (
	// Applies prefix "Warning:" to the log file
	Warn *log.Logger
	// Applies prefix "Info:" to the log file
	Info *log.Logger
	// Applies prefix "Error:" to the log file
	Error *log.Logger
)

func Create() {
	logs, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not create/open a log file", err)
	}

	flags := log.LstdFlags | log.Lshortfile

	Info = log.New(logs, "INFO:", flags)
	Warn = log.New(logs, "WARNING:", flags)
	Error = log.New(logs, "ERROR:", flags)

	log.SetOutput(logs)
}
