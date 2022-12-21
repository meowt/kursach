package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogInit checks existing today's log file and if it's true, it sets output to this file.
// Other way - it creates today's file.
func LogInit() (err error) {
	t := time.Now()
	logFileName := fmt.Sprint("logs/" + t.Format("2006-01-02") + ".log")
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log file opening error: %v", err)
	}
	log.SetOutput(file)
	log.Println("---LOGGING-STARTED----------------------------------------------------------------------------------------------------------")
	return
}
