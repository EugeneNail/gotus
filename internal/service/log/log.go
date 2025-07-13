package log

import (
	"fmt"
	"github.com/EugeneNail/gotus/internal/enum/environment"
	"io"
	goLog "log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var today int = -1

var infoLogger *goLog.Logger
var debugLogger *goLog.Logger
var errorLogger *goLog.Logger
var file *os.File

func Initialize() {
	ticker := time.NewTicker(time.Second)
	setLoggers()

	go func() {
		for range ticker.C {
			if today != time.Now().Day() {
				setLoggers()
				today = time.Now().Day()
			}
		}
	}()

}

func setLoggers() {
	if file != nil {
		if err := file.Close(); err != nil {
			panic(fmt.Errorf("cannot close the log file: %w", err))
		}
	}

	file = openFile()
	infoLogger = goLog.New(io.MultiWriter(file, os.Stdout), "[ INFO  ] ", goLog.Ldate|goLog.Ltime|goLog.Lmicroseconds|goLog.Lmsgprefix)
	debugLogger = goLog.New(io.MultiWriter(file, os.Stdout), "[ DEBUG ] ", goLog.Ldate|goLog.Ltime|goLog.Lmicroseconds|goLog.Lmsgprefix)
	errorLogger = goLog.New(io.MultiWriter(file, os.Stderr), "[ ERROR ] ", goLog.Ldate|goLog.Ltime|goLog.Lmicroseconds|goLog.Lmsgprefix)
}

func openFile() *os.File {
	filename := path.Join(os.Getenv("APP_ROOT"), "storage", "logs", time.Now().Format("2006-01-02.log"))
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("cannot create a log file: %w", err))
	}

	return file
}

func RedirectPanicToLogger() {
	if x := recover(); x != nil {
		Error(x)
		os.Exit(1)
	}
}

// Info writes a message prefixed with [ INFO  ] into a file and Stdout
func Info(messages ...any) {
	infoLogger.Println(messages...)
}

// Debug writes a message prefixed with [ DEBUG ] into a file and Stdout and including file name and line number, but only if ENVIRONMENT is "development"
func Debug(messages ...any) {
	if os.Getenv("ENVIRONMENT") == environment.Development.ToString() {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic(fmt.Errorf("cannot recover caller information"))
		}

		location := strings.ReplaceAll(fmt.Sprintf("at %s:%d", file, line), os.Getenv("APP_ROOT"), "")
		messages = append(messages, location)
		debugLogger.Println(messages...)
	}
}

// Error writes a message prefixed with [ ERROR  ] into a file and Stderr
func Error(messages ...any) {
	errorLogger.Println(messages...)
}
