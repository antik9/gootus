package events

import (
	"fmt"
	"io"
	"time"
)

type HwAccepted struct {
	Id    int
	Grade int
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

type OtusEvent interface {
	ToLogString() string
}

func (h HwAccepted) ToLogString() string {
	t := time.Now()
	return fmt.Sprintf(
		"%d-%02d-%02d accepted %d %d\n", t.Year(), t.Month(), t.Day(), h.Id, h.Grade)
}

func (h HwSubmitted) ToLogString() string {
	t := time.Now()
	return fmt.Sprintf(
		"%d-%02d-%02d submitted %d \"%s\"\n", t.Year(), t.Month(), t.Day(), h.Id, h.Comment)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	w.Write([]byte(e.ToLogString()))
}
