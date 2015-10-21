package logsetup

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type LogRecord struct {
	ct           time.Time
	Name         string
	Message      string
	LevelNo      string
	LevelName    string
	PathName     string
	FileName     string
	Package      string
	FuncName     string
	LineNo       int
	Args         []interface{}
	Created      float64
	MilliSeconds int
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
	r.Args = args
	r.ct = time.Now()
	r.Created = float64(r.ct.Local().UnixNano()/1000/1000) / 1000
	r.MilliSeconds = r.ct.Nanosecond() / 1000 / 1000
	return r
}
func (r *LogRecord) GetMessage() string {
	msg := r.Message
	if r.Args != nil {
		msg = fmt.Sprintf(msg, r.Args...)
	}
	return msg
}
func (r *LogRecord) String() string {
	return fmt.Sprintf("<LogRecord: %v, %v, %v, %v, \"%v\">",
		r.Name, r.LevelNo, r.PathName, r.LineNo, r.Message)
}
func (r *LogRecord) GetRuntimeInfo() {
	pc, pathname, lineno, ok := runtime.Caller(0)
	if !ok {
		r.PathName = "???"
		r.FileName = "???"
		r.FuncName = "???"
		r.LineNo = 0
	} else {
		r.PathName = pathname
		r.LineNo = lineno
		r.FileName = path.Base(PathName)
		r.FuncName = runtime.FuncForPC(pc).Name()
	}
}
