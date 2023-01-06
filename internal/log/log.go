package log

import (
	"os"
	// "bytes"
	"fmt"
	golog "log"
	"sync"
)

var lock = &sync.Mutex{}

var loggerInstance *golog.Logger

/* returns a singleton of the built-in `log.Logger` instance */
func getStdOutLogger() *golog.Logger {
	// fmt.Println("looking for instance")
	if loggerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if loggerInstance == nil {
			prefix := ""
			// 2023/01/06 06:28:39.751550 log.go:49: debug: debug?
			loggerInstance = golog.New(os.Stdout, prefix, golog.Ldate|golog.Ltime|golog.Lmicroseconds|golog.LUTC|golog.Lshortfile)
		}
	}
	return loggerInstance
}

type ILog interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

/* simplest implementation of logging, writing to stdout with a level prefix */

type StdOutLogger struct {
	ILog
}

func write(loggerInstance *golog.Logger, prefix string, msg string) {
	// "warn: the CPU temperature has reached boiling point"
	callDepth := 4 // log.write, type dispatch Debug, convenience Debug, *actual interesting fn*
	loggerInstance.Output(4, fmt.Sprintf("%s: %s\n", prefix, msg))
}

func (l StdOutLogger) Debug(msg string) {
	write(getStdOutLogger(), "debug", msg)
}

func (l StdOutLogger) Info(msg string) {
	write(getStdOutLogger(), "info", msg)
}

func (l StdOutLogger) Warn(msg string) {
	write(getStdOutLogger(), "warn", msg)
}

func (l StdOutLogger) Error(msg string) {
	write(getStdOutLogger(), "error", msg)
}

/* convenience. apps can just go `log.Info("some info message")` at
the expense of controlling how the message is formatted and where it goes.
by default log messages use the built-in `log` package and are written to stdout.
*/

/* inspects application state and returns a logger instance. */
func configuredLogger() ILog {
	var logger_inst ILog
	if true {
		logger_inst = StdOutLogger{}
	}
	return logger_inst
}

func Debug(msg string) {
	configuredLogger().Debug(msg)
}

func Info(msg string) {
	configuredLogger().Info(msg)
}

func Warn(msg string) {
	configuredLogger().Warn(msg)
}

func Error(msg string) {
	configuredLogger().Error(msg)
}
