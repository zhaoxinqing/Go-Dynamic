package logger

import (
	"demo/public"
	"log"
	"os"
	"time"
)

var Logger MyLogger

const (
	LOG_PREFIX_INFO  = "[INFO] "
	LOG_PREFIX_WARN  = "[WARN] "
	LOG_PREFIX_ERROR = "[ERROR] "
)

type MyLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func NewMyLogger() *MyLogger {
	return &MyLogger{
		infoLogger:  log.New(os.Stdout, "", 0),
		warnLogger:  log.New(os.Stdout, "", 0),
		errorLogger: log.New(os.Stderr, "", log.Lshortfile),
	}
	// Logger = MyLogger{
	// 	infoLogger:  log.New(os.Stdout, "", 0),
	// 	warnLogger:  log.New(os.Stdout, "", 0),
	// 	errorLogger: log.New(os.Stderr, "", log.Lshortfile),
	// }
}

func (l *MyLogger) Info(v ...interface{}) {
	l.infoLogger.SetPrefix(LOG_PREFIX_INFO + time.Now().Format(public.TIME_FORMAT) + " : ")
	l.infoLogger.Println(v...)
}

func (l *MyLogger) Warn(v ...interface{}) {
	l.warnLogger.SetPrefix(LOG_PREFIX_WARN + time.Now().Format(public.TIME_FORMAT) + " : ")
	l.warnLogger.Println(v...)
}

func (l *MyLogger) Error(v ...interface{}) {
	l.errorLogger.SetPrefix(LOG_PREFIX_ERROR + time.Now().Format(public.TIME_FORMAT) + " : ")
	l.errorLogger.Println(v...)
}
