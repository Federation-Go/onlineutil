package logsetup

import (
	"strings"
)

type Filter string

func NewFilter(name string) Filter {
	return Filter(name)
}
func (f Filter) Filter(record *LogRecord) bool {
	length := len(f)
	switch {
	case length == 0:
		return true
	case string(f) == record.Name:
		return true
	case strings.Index(record.Name, string(f)) != 0:
		return false
	}
	return string(record.Name[length]) == "."
}
