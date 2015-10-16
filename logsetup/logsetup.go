package logsetup

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

const (
	CRITICAL = 50
	FATAL    = CRITICAL
	ERROR    = 40
	WARNING  = 30
	WARN     = WARNING
	INFO     = 20
	DEBUG    = 10
	NOTSET   = 0
)

var levelNames = map[string]int{
	"CRITICAL": CRITICAL,
	"ERROR":    ERROR,
	"WARN":     WARNING,
	"WARNING":  WARNING,
	"INFO":     INFO,
	"DEBUG":    DEBUG,
	"NOTSET":   NOTSET,
}
var lock = sync.Mutex

func acquireLock() {
	lock.Lock()
}
func releaseLock() {
	lock.Unlock()
}
func checkLevel(level string) error {
	if value, ok := levelNames[level]; ok {
		return nil
	}
	return errors.New(fmt.Printf("Unknown level: %s", level))
}

func addLevelName(level int, levelName string) {
	acquireLock()
	defer releaseLock()
	levelNames[levelName] = level
}

type Logger struct {
	Name      string
	Level     string
	Parent    Logger
	Propagate bool
	Handlers  []Handler
	Disable   bool
}

func (logger *Logger) SetLevel(level string) error {
	err := checkLevel(level)
	if err != nil {
		return err
	}
	logger.Level = level
	return nil
}
