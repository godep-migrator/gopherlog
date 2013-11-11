package gopherlog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	ROOT_LOGGER_NAME = "<root>"
	HOSTNAME_UNKNOWN = "unknown"
)

type levelHandler struct {
	h Handler
	l Level
}

var (
	defaultHandler = &IOHandler{Out: os.Stderr}
	handlers       = make([]levelHandler, 0, 2)
	mu             sync.RWMutex
)

func RegisterHandler(h Handler, l Level) {
	mu.Lock()
	defer mu.Unlock()

	handlers = append(handlers, levelHandler{h: h, l: l})
}

func ClearHandlers() {
	mu.Lock()
	defer mu.Unlock()

	handlers = make([]levelHandler, 0)
}

func getBaseData() (out map[string]interface{}) {
	var (
		err      error
		hostname string
	)

	out = make(map[string]interface{})
	out["pid"] = os.Getpid()
	if hostname, err = os.Hostname(); err != nil {
		hostname = HOSTNAME_UNKNOWN
	}
	out["hostname"] = hostname
	out["time"] = time.Now()
	out["name"] = ROOT_LOGGER_NAME

	return
}

func addBaseData(data map[string]interface{}) (out map[string]interface{}) {
	var (
		baseData = getBaseData()
	)

	if data == nil {
		return baseData
	} else {
		if _, hasname := data["name"]; hasname {
			delete(baseData, "name")
		}
		out = data
		for k, v := range baseData {
			out[k] = v
		}
	}
	return
}

func Log(l Level, v ...interface{}) error {
	return Logm(l, fmt.Sprint(v...), nil)
}

func Logm(l Level, message string, data map[string]interface{}) (err error) {
	data = addBaseData(data)

	mu.RLock()
	defer mu.RUnlock()

	if len(handlers) == 0 {
		err = defaultHandler.Log(l, message, data)
		return
	}

	for _, handler := range handlers {
		if l >= handler.l {
			err = handler.h.Log(l, message, data)
		}
	}

	return
}

func Debug(v ...interface{}) error {
	return Log(DEBUG, v...)
}

func Info(v ...interface{}) error {
	return Log(INFO, v...)
}

func Warning(v ...interface{}) error {
	return Log(WARNING, v...)
}

func Error(v ...interface{}) error {
	return Log(ERROR, v...)
}

func Critical(v ...interface{}) error {
	return Log(CRITICAL, v...)
}

func Fatal(v ...interface{}) (err error) {
	err = Log(FATAL, v...)
	os.Exit(1)
	return
}

func Debugf(s string, v ...interface{}) error {
	return Log(DEBUG, fmt.Sprintf(s, v...))
}

func Infof(s string, v ...interface{}) error {
	return Log(INFO, fmt.Sprintf(s, v...))
}

func Warningf(s string, v ...interface{}) error {
	return Log(WARNING, fmt.Sprintf(s, v...))
}

func Errorf(s string, v ...interface{}) error {
	return Log(ERROR, fmt.Sprintf(s, v...))
}

func Criticalf(s string, v ...interface{}) error {
	return Log(CRITICAL, fmt.Sprintf(s, v...))
}

func Fatalf(s string, v ...interface{}) (err error) {
	err = Log(FATAL, fmt.Sprintf(s, v...))
	os.Exit(1)
	return
}

func Debugm(s string, m map[string]interface{}) error {
	return Logm(DEBUG, s, m)
}

func Infom(s string, m map[string]interface{}) error {
	return Logm(INFO, s, m)
}

func Warningm(s string, m map[string]interface{}) error {
	return Logm(WARNING, s, m)
}

func Errorm(s string, m map[string]interface{}) error {
	return Logm(ERROR, s, m)
}

func Criticalm(s string, m map[string]interface{}) error {
	return Logm(CRITICAL, s, m)
}

func Fatalm(s string, m map[string]interface{}) (err error) {
	err = Logm(FATAL, s, m)
	os.Exit(1)
	return
}

// ========================================================
// ===================== Named Loggers ====================
// ========================================================

type Logger struct {
	Name string
}

func GetLogger(name string) *Logger {
	return &Logger{Name: name}
}

func (lo *Logger) Log(l Level, v ...interface{}) error {
	return lo.Logm(l, fmt.Sprint(v...), nil)
}

func (lo *Logger) Logm(l Level, message string, data map[string]interface{}) error {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["name"] = lo.Name
	return Logm(l, message, data)
}

func (lo *Logger) Debug(v ...interface{}) error {
	return lo.Log(DEBUG, v...)
}

func (lo *Logger) Info(v ...interface{}) error {
	return lo.Log(INFO, v...)
}

func (lo *Logger) Warning(v ...interface{}) error {
	return lo.Log(WARNING, v...)
}

func (lo *Logger) Error(v ...interface{}) error {
	return lo.Log(ERROR, v...)
}

func (lo *Logger) Critical(v ...interface{}) error {
	return lo.Log(CRITICAL, v...)
}

func (lo *Logger) Fatal(v ...interface{}) (err error) {
	err = lo.Log(FATAL, v...)
	os.Exit(1)
	return
}

func (lo *Logger) Debugf(s string, v ...interface{}) error {
	return lo.Log(DEBUG, fmt.Sprintf(s, v...))
}

func (lo *Logger) Infof(s string, v ...interface{}) error {
	return lo.Log(INFO, fmt.Sprintf(s, v...))
}

func (lo *Logger) Warningf(s string, v ...interface{}) error {
	return lo.Log(WARNING, fmt.Sprintf(s, v...))
}

func (lo *Logger) Errorf(s string, v ...interface{}) error {
	return lo.Log(ERROR, fmt.Sprintf(s, v...))
}

func (lo *Logger) Criticalf(s string, v ...interface{}) error {
	return lo.Log(CRITICAL, fmt.Sprintf(s, v...))
}

func (lo *Logger) Fatalf(s string, v ...interface{}) (err error) {
	err = lo.Log(FATAL, fmt.Sprintf(s, v...))
	os.Exit(1)
	return
}

func (lo *Logger) Debugm(s string, m map[string]interface{}) error {
	return lo.Logm(DEBUG, s, m)
}

func (lo *Logger) Infom(s string, m map[string]interface{}) error {
	return lo.Logm(INFO, s, m)
}

func (lo *Logger) Warningm(s string, m map[string]interface{}) error {
	return lo.Logm(WARNING, s, m)
}

func (lo *Logger) Errorm(s string, m map[string]interface{}) error {
	return lo.Logm(ERROR, s, m)
}

func (lo *Logger) Criticalm(s string, m map[string]interface{}) error {
	return lo.Logm(CRITICAL, s, m)
}

func (lo *Logger) Fatalm(s string, m map[string]interface{}) (err error) {
	err = lo.Logm(FATAL, s, m)
	os.Exit(1)
	return err
}
