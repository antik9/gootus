package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func checkDir(directory string) {
	src, err := os.Stat(directory)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if !src.IsDir() {
		log.Fatalf("%s is not a directory", directory)
	}
}

func exportEnvFromDir(directory string) {
	checkDir(directory)

	src, err := os.Open(directory)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer src.Close()

	fileSlice, err := src.Readdirnames(0)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, filename := range fileSlice {
		exportEnvFromFile(filepath.Join(src.Name(), filename))
	}
}

func exportEnvFromFile(filename string) {
	envName := filepath.Base(filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}

	os.Setenv(envName, strings.Trim(string(content), " \n\t\r\b\x00"))
}

func help() {
	fmt.Println(`Usage:
	gootus <env_dir> <command> ...args`)
	os.Exit(0)
}

func execute(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
