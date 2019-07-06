package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPlainCopy(t *testing.T) {
	REFERENCE := "rather simple content"
	content := []byte(REFERENCE)
	result := make([]byte, 1024)

	src, errSrc := ioutil.TempFile(".", "testSrc*")
	failOnErr(errSrc)
	defer src.Close()
	defer os.Remove(src.Name())

	dest, errDest := ioutil.TempFile(".", "testDest*")
	failOnErr(errDest)
	dest.Close()
	defer os.Remove(dest.Name())

	writer := os.NewFile(3, src.Name())
	defer writer.Close()
	writer.Write(content)

	Copy(src.Name(), dest.Name(), 0, 1024)

	reader, err := os.Open(dest.Name())
	failOnErr(err)
	defer reader.Close()
	reader.Read(result)

	if string(content) != strings.Trim(string(result), "\x00") {
		t.Errorf("get %s; want %s", string(result), string(content))
	}
}

func TestCopyWithOffset(t *testing.T) {
	REFERENCE := "rather simple content"
	content := []byte(REFERENCE)
	result := make([]byte, 1024)

	src, errSrc := ioutil.TempFile(".", "testSrc*")
	failOnErr(errSrc)
	defer src.Close()
	defer os.Remove(src.Name())

	dest, errDest := ioutil.TempFile(".", "testDest*")
	failOnErr(errDest)
	dest.Close()
	defer os.Remove(dest.Name())

	writer := os.NewFile(3, src.Name())
	defer writer.Close()
	writer.Write(content)

	Copy(src.Name(), dest.Name(), 10, 1024)

	reader, err := os.Open(dest.Name())
	failOnErr(err)
	defer reader.Close()
	reader.Read(result)

	if string(content[10:]) != strings.Trim(string(result), "\x00") {
		t.Errorf("get %s; want %s", string(result), string(content[10:]))
	}
}

func TestCopyWithLimit(t *testing.T) {
	REFERENCE := "rather simple content"
	content := []byte(REFERENCE)
	result := make([]byte, 1024)

	src, errSrc := ioutil.TempFile(".", "testSrc*")
	failOnErr(errSrc)
	defer src.Close()
	defer os.Remove(src.Name())

	dest, errDest := ioutil.TempFile(".", "testDest*")
	failOnErr(errDest)
	dest.Close()
	defer os.Remove(dest.Name())

	writer := os.NewFile(3, src.Name())
	defer writer.Close()
	writer.Write(content)

	Copy(src.Name(), dest.Name(), 0, 10)

	reader, err := os.Open(dest.Name())
	failOnErr(err)
	defer reader.Close()
	reader.Read(result)

	if string(content[:10]) != strings.Trim(string(result), "\x00") {
		t.Errorf("get %s; want %s", string(result), string(content[:10]))
	}
}
