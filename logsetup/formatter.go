package logsetup

var formats = map[string]string{
	"(name)":      "name",
	"(time)":      "time",
	"(levelname)": "levelname",
	"(funcname)":  "funcname",
	"(message)":   "message",
	"(longfile)":  "shortfile",
	"(utc)":       "utc",
}

const (
	Ltime = 1 << itoa
	Lname
	Llevelname
	Lfuncname
	Lmessage
	Llongfile
	Lshortfile
	LUTC
	LstdFlags = Ldate | Ltime | LUTC | Lmessage
)

type Formatter struct {
	LogFormat  string
	TimeLayout string
}

var defaultLayout = "2006-01-02 15:04:05.999"
var defaultFormat = "%(message)"

func NewFormatter(logFormat string, timeLayout string) *Formatter {
	var f = new(Formatter)
	if f.LogFormat = logFormat; logFormat == "" {
		f.LogFormat = LstdFlags
	}
	if f.TimeLayout = timeLayout; timeLayout == "" {
		f.TimeLayout = defaultLayout
	}
	return f
}
func (f Formatter) FormatTime(record LogRecord, layout string) string {
	var date string
	if layout != "" {
		date = record.Created.Format(layout)
	} else {
		date = record.Created.Format(defaultLayout)
	}
	return date
}
func (f Formatter) Format(record LogRecord) string {
	message := record.GetMessage()
}
