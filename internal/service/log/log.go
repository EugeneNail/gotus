package log

import (
	"fmt"
	"github.com/EugeneNail/gotus/internal/enum/environment"
	"io"
	goLog "log"
	"os"
	"path"
	"time"
)

var today int = -1

var infoLogger *goLog.Logger
var debugLogger *goLog.Logger
var errorLogger *goLog.Logger

func Initialize() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C

		if today != time.Now().Day() {
			filename := path.Join(os.Getenv("APP_ROOT"), "storage", "logs", time.Now().Format("2006-01-02.log"))

			file, err := os.Create(filename)
			if err != nil {
				panic(fmt.Errorf("cannot create a log file: %w", err))
			}

			infoLogger = goLog.New(io.MultiWriter(file, os.Stdout), "[ INFO  ] ", goLog.Ldate|goLog.Ltime|goLog.Lmsgprefix)
			debugLogger = goLog.New(io.MultiWriter(file, os.Stdout), "[ DEBUG ] ", goLog.Ldate|goLog.Ltime|goLog.Lmsgprefix|goLog.Lshortfile)
			errorLogger = goLog.New(io.MultiWriter(file, os.Stderr), "[ ERROR ] ", goLog.Ldate|goLog.Ltime|goLog.Lmsgprefix)

			today = time.Now().Day()
		}
	}
}

// Info writes a message prefixed with [ INFO  ] into a file and Stdout
func Info(messages ...any) {
	infoLogger.Println(messages)
}

// Debug writes a message prefixed with [ DEBUG ] into a file and Stdout, but only if ENVIRONMENT is "development"
func Debug(messages ...any) {
	if os.Getenv("ENVIRONMENT") == environment.Development.ToString() {
		debugLogger.Println(messages)
	}
}

// Error writes a message prefixed with [ ERROR  ] into a file and Stderr
func Error(messages ...any) {
	errorLogger.Println(messages)
}
