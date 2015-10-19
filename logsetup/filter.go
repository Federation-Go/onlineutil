package logsetup

import (
	"strings"
)

type Filter string

func NewFilter(name string) Filter {
	var f = Filter(name)
	return f
}
func (f Filter) filter(record LogRecord) bool {
	length := len(f)
	switch {
	case length == 0:
		return true
	case f == record.Name:
		return true
	case strings.Index(record.Name, f) != 0:
		return false
	}
	return string(record.Name[length]) == "."
}
