package gopherlog

import (
	"encoding/json"
	"fmt"
	"github.com/kisielk/raven-go/raven"
	"io"
	"strings"
	"time"
)

const (
	BUNYAN_TIME_FORMAT = "2006-01-02T15:04:05-07:00"
	SENTRY_TIME_FORMAT = "2006-01-02T15:04:05"
)

var (
	bunyanLevel = map[Level]int{
		DEBUG:    20,
		INFO:     30,
		WARNING:  40,
		ERROR:    50,
		CRITICAL: 60,
		FATAL:    60,
	}
)

type Handler interface {
	// Log writes the message to the desired output stream
	// returns error if it can't log.
	Log(l Level, message string, data map[string]interface{}) error
}

func copyData(in map[string]interface{}) (out map[string]interface{}) {
	out = make(map[string]interface{})
	for k, v := range in {
		out[k] = v
	}
	return
}

// IOHandler logs everything to the supplied io.Writer.
type IOHandler struct {
	Out io.Writer
}

func (i *IOHandler) Log(l Level, message string, data map[string]interface{}) error {
	var (
		accumulator = make([]string, 0, 10)
		err         error
		time        = data["time"].(time.Time).Format(SENTRY_TIME_FORMAT)
	)

	for k, v := range data {
		accumulator = append(accumulator, fmt.Sprintf("%s = %s", k, v))
	}

	message = message + " " + strings.Join(accumulator, " ")

	_, err = fmt.Fprintf(i.Out, "(%s) [%s] %s: %s\n", data["hostname"], time, l, message)
	return err
}

type BunyanHandler struct {
	Out io.Writer
}

func (b *BunyanHandler) Log(l Level, message string, data map[string]interface{}) error {
	outData := copyData(data)

	time := data["time"].(time.Time).Format(BUNYAN_TIME_FORMAT)
	outData["time"] = time
	outData["v"] = 0
	outData["level"] = bunyanLevel[l]
	outData["msg"] = message
	enc := json.NewEncoder(b.Out)
	return enc.Encode(outData)
}

type RavenHandler struct {
	DSN         string
	ProjectName string
	client      *raven.Client
}

func NewRavenHandler(projectName, dsn string) *RavenHandler {
	return &RavenHandler{DSN: dsn, ProjectName: projectName}
}

func (r *RavenHandler) dataToEvent(l Level, message string, data map[string]interface{}) *raven.Event {
	//TODO: fix date formatting
	//TODO: figure out how to add other fields
	//EventId is left as default. It will be filled by go-raven
	ev := &raven.Event{
		Project:   r.ProjectName,
		Message:   message,
		Timestamp: data["time"].(time.Time).Format(SENTRY_TIME_FORMAT),
		Level:     l.String(),
		Logger:    data["name"].(string),
	}

	return ev
}

func (r *RavenHandler) Log(l Level, message string, data map[string]interface{}) error {
	if r.client == nil {
		if client, err := raven.NewClient(r.DSN); err != nil {
			return err
		} else {
			r.client = client
		}
	}

	ev := r.dataToEvent(l, message, data)
	err := r.client.Capture(ev)
	return err
}
