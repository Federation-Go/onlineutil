package logsetup

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type LogRecord struct {
	Name      string
	Message   string
	LevelNo   string
	LevelName string
	PathName  string
	LineNo    int
	Args      []interface{}
	Created   time.Time
}

func NewLogRecord(name, message, levelName string, args ...interface{}) *LogRecord {
	var r = new(LogRecord)
	r.Name = name
	r.Message = message
	levelNo, err := checkLevel(levelName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create LogRecord, unknown level name %v\n", levelName)
		os.Exit(1)
	}
	r.LevelName = levelName
	r.LevelNo = levelNo
	var ok bool
	_, r.PathName, r.LineNo, ok = runtime.Caller(0)
	if !ok {
		r.PathName = "???"
		r.LineNo = 0
	}
	r.Args = args
	r.Created = time.Now()
	return r
}
func (r LogRecord) GetMessage() string {
	msg := r.Message
	if r.Args != nil {
		msg = fmt.Sprintf(msg, r.Args...)
	}
	return msg
}
func (r LogRecord) String() string {
	return fmt.Sprintf("<LogRecord: %v, %v, %v, %v, \"%v\">",
		r.Name, r.LevelNo, r.PathName, r.LineNo, r.Message)
}
