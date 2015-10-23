package logsetup

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type DefaultHandler struct {
	Filterer
	name      string
	LevelName string
	Formatter *Formatter
	Lock      sync.Mutex
}

func NewDefaultHandler(name string, levelname string, formatter *Formatter) *Handler {
	var h = new(DefaultHandler)
	return h
}
func (h *Handler) Name() string {
	return h.name
}
func (h *Handler) SetName(name string) {
	acquireLock()
	defer releaseLock()

}
func (h *Handler) Acquire() {
	h.Lock.Lock()
}
func (h *Handler) Release() {
	h.Lock.Unlock()
}
func (h *Handler) setLevel(levelname string) error {
	_, err := checkLevel(levelname)
	if err != nil {
		return err
	}
	h.LevelName = levelname
	return nil
}
func (h *Handler) LevelNo() int {
	levelno, err := checkLevel(h.LevelName)
	if err != nil {
		fmt.Sprint(os.Stderr, "handler's levelname is invalid\n")
		os.Exit(1)
	}
	return levelno
}
func (h *Handler) Format(record *LogRecord) string {
	if h.Formatter == nil {
		h.Formatter = DefaultFormatter
	}
	return h.Formatter.Format(record)
}
func (h *Handler) Emit(record LogRecord) error {
	return errors.New("emit function not implements by handler")
}
func (h *Handler) Handle(record *LogRecord) {
	rv := h.Filter(record)
	if rv {
		h.Acquire()
		defer h.Release()

	}
}
