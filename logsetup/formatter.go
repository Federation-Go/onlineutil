package logsetup

import (
	"strconv"
	"strings"
)

var formats = map[string]string{
	"%(name)s":      "name",
	"%(levelno)s":   "levelno",
	"%(levelname)s": "levelname",
	"%(pathname)s":  "pathname",
	"%(filename)s":  "filename",
	"%(package)s":   "package",
	"%(lineno)d":    "lineno",
	"%(funcName)s":  "funcName",
	"%(created)f":   "created",
	"%(asctime)s":   "asctime",
	"%(msecs)d":     "msecs",
	"%(message)s":   "message",
}

type Formatter struct {
	LogFormat  string
	TimeLayout string
}

var defaultTimeLayout = "2006-01-02 15:04:05,999"
var defaultFormat = "%(message)s"

func NewFormatter(logFormat string, timeLayout string) *Formatter {
	var f = new(Formatter)
	if f.LogFormat = logFormat; logFormat == "" {
		f.LogFormat = defaultFormat
	}
	if f.TimeLayout = timeLayout; timeLayout == "" {
		f.TimeLayout = defaultTimeLayout
	}
	return f
}
func (f *Formatter) FormatTime(record *LogRecord, layout string) string {
	var date string
	if layout != "" {
		date = record.ct.Format(layout)
	} else {
		date = record.ct.Format(defaultTimeLayout)
	}
	return date
}
func (f *Formatter) Format(record *LogRecord) string {
	message := record.GetMessage()
	var asctime string
	if f.UseTime() {
		asctime = f.FormatTime(record, f.TimeLayout)
	}
	log := f.LogFormat
	for index, key := range strings.Split(f.LogFormat, " ") {
		if value, ok := formats[key]; ok {
			switch value {
			case "name":
				log = strings.Replace(f.LogFormat, key, record.Name, -1)
			case "levelno":
				log = strings.Replace(f.LogFormat, key, strconv.Itoa(record.LevelNo), -1)
			case "levelname":
				log = strings.Replace(f.LogFormat, key, record.LevelName, -1)
			case "pathname":
				log = strings.Replace(f.LogFormat, key, record.PathName, -1)
			case "filename":
				log = strings.Replace(f.LogFormat, key, record.FileName, -1)
			case "lineno":
				log = strings.Replace(f.LogFormat, key, strconv.Itoa(record.LineNo), -1)
			case "package":
				log = strings.Replace(f.LogFormat, key, record.PackageName, -1)
			case "created":
				log = strings.Replace(f.LogFormat, key,
					strconv.FormatFloat(record.Created, 'f', 3, 64), -1)
			case "asctime":
				log = strings.Replace(f.LogFormat, key, asctime, -1)
			case "msecs":
				log = strings.Replace(f.LogFormat, key, strconv.Itoa(record.MilliSeconds), -1)
			case "message":
				log = strings.Replace(f.LogFormat, key, message, -1)
			}
		}
	}
	return log
}
func (f *Formatter) UseTime() bool {
	if index := strings.Index(f.LogFormat, "%(asctime)"); index >= 0 {
		return true
	}
	return false
}
