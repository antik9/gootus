package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestExportEnvFromFile(t *testing.T) {
	src, err := ioutil.TempFile(".", "VAR*")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	defer src.Close()
	defer os.Remove(src.Name())

	writer := os.NewFile(3, src.Name())
	defer writer.Close()
	if _, err := writer.Write([]byte("XOX")); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	exportEnvFromFile(src.Name())

	if os.Getenv(src.Name()) != "XOX" {
		t.Errorf("get %s; want XOX", os.Getenv(src.Name()))
	}
}

func TestExportEnvFromDir(t *testing.T) {
	tmpDir, err := ioutil.TempDir(".", "envs")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	defer os.RemoveAll(tmpDir)

	src1, err := ioutil.TempFile(tmpDir, "VAR*")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	defer src1.Close()

	src2, err := ioutil.TempFile(tmpDir, "VAR*")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	defer src2.Close()

	src3, err := ioutil.TempFile(tmpDir, "VAR*")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	defer src3.Close()

	writer1 := os.NewFile(3, src1.Name())
	defer writer1.Close()
	if _, err = writer1.WriteString(src1.Name()); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	writer2 := os.NewFile(5, src2.Name())
	defer writer2.Close()
	if _, err = writer2.WriteString(src2.Name()); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	writer3 := os.NewFile(6, src3.Name())
	defer writer3.Close()
	if _, err = writer3.WriteString(src3.Name()); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	exportEnvFromDir(tmpDir)

	for _, file := range []*os.File{src1, src2, src3} {
		if os.Getenv(filepath.Base(file.Name())) != file.Name() {
			t.Errorf("get %s; want %s", os.Getenv(file.Name()), file.Name())
		}
	}
}
