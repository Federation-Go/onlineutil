package logsetup

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

const (
	Ldate = 1 << itoa
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	Lname
	Llevelname
	Llevelno
	Lfuncname
	Lmessage
	LstdFlags = Ldate | Ltime | LUTC | Lmessage
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

var DefaultFormatter = new(Formatter)

type LogRecord struct {
	Name      string
	Message   string
	LevelNo   string
	LevelName string
	PathName  string
	LineNo    string
	Args      []string
	Created   time.Time
}

func (r LogRecord) GetMessage() string {
	msg := r.Message
	if r.Args != nil {
		msg = fmt.Printf(msg, r.Args)
	}
	return msg
}
func (r LogRecord) String() string {
	return fmt.Sprintf("<LogRecord: %v, %v, %v, %v, \"%v\">",
		r.Name, r.LevelNo, r.PathName, r.LineNo, r.Message)
}
func NewLogRecord(name, levelname, pathname, lineno, message string, args []string) *LogRecord {
	var r = new(LogRecord)
	r.Created = time.Now()
	r.Message = message
	r.PathName = pathname
	r.LineNo = lineno
	r.Name = name
	r.LevelName = levelname
	r.Args = args
	return r
}

type Formatter struct {
	fmt    string
	layout string
}

var defaultLayout = "2006-01-02 15:04:05.999"

func (f Formatter) FormatTime(record LogRecord, layout string) string {
	var date string
	if layout != "" {
		date = record.Created.Format(layout)
	} else {
		date = record.Created.Format(defaultLayout)
	}
	return date
}
func (f Formatter) FormatException() {
}
func (f Formatter) Format(record LogRecord) {
	record.Message = record.getMessage()
}
func NewFormatter(fmt, layout string) *Formatter {
	var formatter = new(Formatter)
	if fmt != nil {
		formatter.fmt = fmt
	} else {
		formatter.fmt = "%(message)s"
	}
	formatter.layout = layout
	return formatter
}
