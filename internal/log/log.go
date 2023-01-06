package log

import (
	"os"
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
	Debug(msg string, a ...any)
	Info(msg string, a ...any)
	Warn(msg string, a ...any)
	Error(msg string, a ...any)
}

/* simplest implementation of logging, writing to stdout with a level prefix */

type StdOutLogger struct {
	ILog
}

func write(loggerInstance *golog.Logger, prefix string, msg string, a ...any) {
	// "warn: the CPU temperature has reached boiling point"
	callDepth := 4 // log.write, type dispatch Debug, convenience Debug, *actual interesting fn*
	rest := fmt.Sprint(a...)
	loggerInstance.Output(callDepth, fmt.Sprintf("%s: %s\n", prefix, msg+rest))
}

func (l StdOutLogger) Debug(msg string, a ...any) {
	write(getStdOutLogger(), "debug", msg, a...)
}

func (l StdOutLogger) Info(msg string, a ...any) {
	write(getStdOutLogger(), "info", msg, a...)
}

func (l StdOutLogger) Warn(msg string, a ...any) {
	write(getStdOutLogger(), "warn", msg, a...)
}

func (l StdOutLogger) Error(msg string, a ...any) {
	write(getStdOutLogger(), "error", msg, a...)
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

func Debug(msg string, a ...any) {
	configuredLogger().Debug(msg, a...)
}

func Info(msg string, a ...any) {
	configuredLogger().Info(msg, a...)
}

func Warn(msg string, a ...any) {
	configuredLogger().Warn(msg, a...)
}

func Error(msg string, a ...any) {
	configuredLogger().Error(msg, a...)
}
