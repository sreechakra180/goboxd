package logger

import (
	"log"
	"os"
)

// keeping it simple for hackathon - a real project would use zap or logrus
// but standard log is fast to setup
var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	WarnLog  *log.Logger
)

func InitLogger() {
	flags := log.Ldate | log.Ltime | log.Lshortfile

	InfoLog = log.New(os.Stdout, "INFO: ", flags)
	WarnLog = log.New(os.Stdout, "WARN: ", flags)
	ErrorLog = log.New(os.Stderr, "ERROR: ", flags)
}
