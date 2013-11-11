package gopherlog

import (
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIOHandler(t *testing.T) {
	var (
		b       bytes.Buffer
		handler = &IOHandler{Out: &b}
	)

	ClearHandlers()

	Convey("Given a IOHandler with a level of WARNING", t, func() {
		RegisterHandler(handler, WARNING)
		Convey("the root logger", func() {
			Convey("should log WARNING, ERROR and CRITICAL messages", func() {
				Warning("warning_message")
				So(b.String(), ShouldContainSubstring, "warning_message")
				b.Reset()
				Error("error_message")
				So(b.String(), ShouldContainSubstring, "error_message")
				b.Reset()
				Warning("critical_message")
				So(b.String(), ShouldContainSubstring, "critical_message")
				Convey("and it should log them as the root logger", func() {
					So(b.String(), ShouldContainSubstring, "name = <root>")
				})
			})
			b.Reset()

			Convey("it shouldn't log DEBUG or INFO message", func() {
				Debug("debug_message")
				So(b.String(), ShouldEqual, "")
				b.Reset()
				Info("info_message")
				So(b.String(), ShouldEqual, "")
			})
			b.Reset()
		})

		testLogger := GetLogger("testLogger")
		Convey("a logger with name 'testLogger'", func() {
			Convey("should log as 'testLogger'", func() {
				testLogger.Error("debug_message")
				So(b.String(), ShouldContainSubstring, "name = testLogger")
				b.Reset()
			})
		})

	})
}

func TestBunyanHandler(t *testing.T) {
	var (
		b       bytes.Buffer
		handler = &BunyanHandler{Out: &b}
	)

	ClearHandlers()

	Convey("Given a BunyanHandler with a level of DEBUG", t, func() {
		RegisterHandler(handler, DEBUG)
		Convey("the output to Log should be valid json", func() {
			Debug("mymessage")
			d := json.NewDecoder(&b)
			value := make(map[string]interface{})
			err := d.Decode(&value)
			So(err, ShouldBeNil)

			Convey("its 'v' field should exist and be equal to 0", func() {
				v, exists := value["v"]
				So(exists, ShouldBeTrue)
				So(v, ShouldEqual, 0)
			})

			Convey("it should have a message", func() {
				_, exists := value["msg"]
				So(exists, ShouldBeTrue)
				Convey("and it should be 'mymessage'", func() {
					So(value["msg"], ShouldEqual, "mymessage")
				})
			})
		})
	})
}
