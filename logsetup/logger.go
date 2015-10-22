package logsetup

import (
	"fmt"
	"os"
	"runtime"
)

type Logger struct {
	*Filterer
	Name      string
	LevelName string
	LevelNo   int
	Parent    *Logger
	Propagate bool
	Handlers  []IHandler
	Disable   bool
}

func NewLogger(name string) *Logger {
	return &Logger{
		Filterer:  NewFilterer(),
		Name:      name,
		LevelName: "NOTSET",
		LevelNo:   0,
		Parent:    nil,
		Propagate: true,
		Handlers:  make([]IHandler, 2, 10),
		Disable:   false,
	}
}
func (l *Logger) SetLevel(levelname string) error {
	levelno, err := checkLevel(levelname)
	if err != nil {
		return err
	}
	l.LevelName = levelname
	l.LevelNo = levelno
	return nil
}
func (l *Logger) Debug(message string, args ...interface{}) {
}
func (l *Logger) Info(message string, args ...interface{}) {
}

func (l *Logger) Warn(message string, args ...interface{}) {
}

func (l *Logger) Error(message string, args ...interface{}) {
}

func (l *Logger) Fatal(message string, args ...interface{}) {
}
func (l *Logger) Log(levelname, message string, args ...interface{}) {
}
func (l *Logger) Handle(r *LogRecord) {
	if !l.Disable && l.Filter(r) {
	}
}
func (l *Logger) CallHandlers(r *LogRecord) {
	i := l
	found := 0
	for i != nil {
		for _, h := range l.Handlers {
			if r.LevelNo > h.LevelNo() {
				h.Handle(r)
				found += 1
			}
		}
		if l.Propagate {
			i = l.Parent
		} else {
			i = nil
		}
	}
	if found == 0 {
		fmt.Fprintf(os.Stderr, "No Handlers could be found for logger \"%s\"\n",
			l.Name)
	}
}
func (l *Logger) AddHandler(h IHandler) {
	acquireLock()
	defer releaseLock()
	for _, v := range l.Handlers {
		if v == h {
			return
		}
	}
	l.Handlers = append(l.Handlers, h)
}
func (l *Logger) RemoveHandler(h IHandler) {
	acquireLock()
	defer releaseLock()
	for i, v := range l.Handlers {
		if v == h {
			l.Handlers = append(l.Handlers[:i], l.Handlers[:i+1]...)
			return
		}
	}

}
func (l *Logger) FindCaller() {
	pc := make([]uintptr, 10)
	count := runtime.Callers(1, pc)
	for i := 0; i < count; i++ {
		f := runtime.FuncForPC(pc[i])

	}
}
