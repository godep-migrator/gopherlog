package gopherlog

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	CRITICAL
	FATAL
)

var logLevels = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"CRITICAL",
	"FATAL",
}

func (l Level) String() string {
	return logLevels[l]
}
