package logsetup

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Filterer struct {
	filters map[Filter]bool
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
func (f *Filterer) filter(record LogRecord) bool {
	for key, _ := range f.filters {
		if !key.filter(record) {
			return false
		}
	}
	return true
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
