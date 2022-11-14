package service

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogInit checks existing today's log file and if it's true, it sets output to this file.
// Other way - it creates today's file.
func LogInit() error {
	t := time.Now()
	logFileName := fmt.Sprint("logs/" + t.Format("02-01-2006") + ".log")
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log file opening error: %v", err)
	}
	//defer file.Close()
	log.SetOutput(file)
	log.Println()
	log.Println("Logging started")
	return err
}
