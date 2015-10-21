package logsetup

import (
	"errors"
	"sync"
)

type IHandler interface {
	Handle(r *LogRecord) bool
}
type Handler struct {
	*Filterer
	name      string
	LevelName string
	Formatter *Formatter
	Lock      sync.Mutex
}

func NewHandler() *Handler {
	var h = new(Handler)
	h.Filterer = NewFilterer()
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
func (h *Handler) setLevel(levelName string) error {
	err := checkLevel(levelName)
	if err != nil {
		return err
	}
	h.LevelName = levelName
	return nil
}
func (h *Handler) Format(record LogRecord) {
	if h.Formatter == nil {
		h.Formatter = DefaultFormatter
	}
	return h.Formatter.Format(record)
}
func (h *Handler) Emit(record LogRecord) error {
	return errors.New("emit function not implements by handler")
}
func (h *Handler) Handle(record LogRecord) {
	rv := h.Filter(record)
	if rv {
		h.Acquire()
		defer h.Release()

	}
}
