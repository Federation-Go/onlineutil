package logsetup

import "strings"

type Manager struct {
	root       Logger
	disable    bool
	loggerDict map[string]*Logger
}

func NewManager(rootnode Logger) *Manager {
	return &Manager{
		root:       rootnode,
		disable:    false,
		loggerDict: make(map[string]*Logger),
	}
}

func (m *Manager) GetLogger(name string) *Logger {
	var l *Logger
	acquireLock()
	defer releaseLock()
	if l, ok := m.loggerDict[name]; !ok {
		l = NewLogger(name)
		m.loggerDict[name] = l
		m.fixupParents(l)
	}
	return l
}
func (m *Manager) fixupParents(l *Logger) {
	name := l.Name
	index := strings.LastIndex(name, ".")
	var parent *Logger
	for index > 0 && parent == nil {
		subname := name[:index]
		if parent, ok := m.loggerDict[subname]; ok {
			break
		}
		index = strings.LastIndex(name[:index], ".")
	}
	if parent != nil {
		l.Parent = parent
	} else {
		l.Parent = root
	}
}
