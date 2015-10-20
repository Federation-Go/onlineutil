package logsetup

import (
	"sync"
)

type Handler struct {
	*Filter
	name      string
	LevelName string
	Formatter *Formatter
	lock      sync.Mutex
}

func (h *Handler) createLock() {

}
func (h *Handler) acquire() {
	h.Lock()
}
func (h *Handler) release() {
	h.Unlock()
}
func (h *Handler) SetLevel(levelName string) {
	err := checkLevel(levelName)
	if err != nil {
		panic("Unknown Level")
	}
	h.LevelName = levelName
}
func (h *Handler) Format(record LogRecord) {
	h.Format(record)
}
func (h *Handler) Handle(record) bool {
	rv := h.filter(record)
}
func (h *Handler) SetFormatter(formatter *Formatter) {
	h.Formatter = formatter
}
