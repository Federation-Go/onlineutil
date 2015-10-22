package logsetup

import (
	"errors"
	"fmt"
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
var lock sync.Mutex

func acquireLock() {
	lock.Lock()
}
func releaseLock() {
	lock.Unlock()
}
func checkLevel(level string) (int, error) {
	if value, ok := levelNames[level]; ok {
		return value, nil
	}
	return -1, errors.New(fmt.Sprintf("Unknown level: %s", level))
}

func addLevelName(level int, levelName string) {
	acquireLock()
	defer releaseLock()
	levelNames[levelName] = level
}

var DefaultFormatter = new(Formatter)

type Filterer struct {
	filters map[Filter]bool
}

func NewFilterer() *Filterer {
	return &Filterer{filters: make(map[Filter]bool)}
}

func (f *Filterer) AddFilter(filter Filter) {
	if _, ok := f.filters[filter]; !ok {
		f.filters[filter] = true
	}
}
func (f *Filterer) RemoveFilter(filter Filter) {
	if _, ok := f.filters[filter]; ok {
		delete(f.filters, filter)
	}
}
func (f *Filterer) Filter(record *LogRecord) bool {
	for key, _ := range f.filters {
		if !key.Filter(record) {
			return false
		}
	}
	return true
}
