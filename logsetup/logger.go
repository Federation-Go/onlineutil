package logsetup

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Logger struct {
	Name      string
	LevelName string
	LevelNo   int
	Parent    *Logger
	Propagate bool
	Handlers  []Handler
	Disable   bool
}

func NewLogger(name string) *Logger {
	return &Logger{
		Name:      name,
		LevelName: "NOTSET",
		LevelNo:   0,
		Parent:    nil,
		Propagate: true,
		Handlers:  make([]Handler, 2, 10),
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
