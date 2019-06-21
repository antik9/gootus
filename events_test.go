package events

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLogStringSubmitted(t *testing.T) {
	now := time.Now()
	reference := fmt.Sprintf(
		"%d-%02d-%02d submitted 5 \"Please, rate my work\"\n",
		now.Year(), now.Month(), now.Day())

	submit := HwSubmitted{5, "", "Please, rate my work"}
	logStr := submit.ToLogString()

	if logStr != reference {
		t.Errorf("get %s; want %s", logStr, reference)
	}
}

func TestLogStringAccepted(t *testing.T) {
	now := time.Now()
	reference := fmt.Sprintf(
		"%d-%02d-%02d accepted 5 3\n",
		now.Year(), now.Month(), now.Day())

	accept := HwAccepted{5, 3}
	logStr := accept.ToLogString()

	if logStr != reference {
		t.Errorf("get %s; want %s", logStr, reference)
	}
}

func TestWriteEvent(t *testing.T) {
	f, err := ioutil.TempFile(".", "test*")

	if err != nil {
		t.Errorf("Problems with test file creation")
	} else {
		writer := os.NewFile(3, f.Name())

		now := time.Now()
		reference := fmt.Sprintf(
			"%d-%02d-%02d submitted 5 \"Please, rate my work\"\n",
			now.Year(), now.Month(), now.Day())

		submit := HwSubmitted{5, "", "Please, rate my work"}
		buffer := make([]byte, len(reference))
		LogOtusEvent(submit, writer)

		writer.Seek(0, 0)
		writer.Read(buffer)

		if string(buffer[:]) != reference {
			t.Errorf("get %s; want %s", string(buffer[:]), reference)
		}

		os.Remove(f.Name())
	}
}
