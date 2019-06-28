package unpacker

import (
	"testing"
)

func TestSimpleString(t *testing.T) {
	result := Unpack("e3io2")
	if result != "eeeioo" {
		t.Errorf("got: %s; want: eeeioo", result)
	}
}

func TestMixWithNumbers(t *testing.T) {
	result := Unpack(`\43i2\12`)
	if result != "444ii11" {
		t.Errorf("got: %s; want: 444ii11", result)
	}
}

func TestEscapes(t *testing.T) {
	result := Unpack(`\\a\\2\3\4\\w1\1`)
	if result != `\a\\34\w1` {
		t.Errorf(`got: %s; want: \a\\34\w1`, result)
	}
}
