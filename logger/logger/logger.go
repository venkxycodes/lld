package logger

import (
	"fmt"
	"sync"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Warn
	Info
	Error
)

func (lvl LogLevel) String() string {
	switch lvl {
	case Debug:
		return "DEBUG"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type Out interface {
	Write(level LogLevel, msg string) error
	Close() error
}

type Logger struct {
	level LogLevel
	out   Out
	mu    sync.Mutex
}

func NewLogger(level LogLevel, sink Out) (*Logger, error) {
	if sink == nil {
		return nil, fmt.Errorf("logger requires a valid out")
	}
	return &Logger{level: level, out: sink}, nil
}

func (l *Logger) log(lvl LogLevel, msg string) {
	if lvl < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	_ = l.out.Write(lvl, msg)
}

func (l *Logger) Debug(msg string) { l.log(Debug, msg) }
func (l *Logger) Warn(msg string)  { l.log(Warn, msg) }
func (l *Logger) Info(msg string)  { l.log(Info, msg) }
func (l *Logger) Error(msg string) { l.log(Error, msg) }

func (l *Logger) Close() error {
	return l.out.Close()
}
