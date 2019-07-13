package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func checkDir(directory string) {
	if src, err := os.Stat(directory); err != nil {
		log.Fatalln(err.Error())
	} else if !src.IsDir() {
		log.Fatalf("%s is not a directory", directory)
	}
}

func exportEnvFromDir(directory string) {
	checkDir(directory)

	if src, err := os.Open(directory); err != nil {
		log.Fatalln(err.Error())
	} else {
		defer src.Close()
		if fileSlice, err := src.Readdirnames(0); err != nil {
			log.Fatalln(err.Error())
		} else {
			for _, filename := range fileSlice {
				exportEnvFromFile(filepath.Join(src.Name(), filename))
			}
		}
	}
}

func exportEnvFromFile(filename string) {
	envName := filepath.Base(filename)
	if src, err := os.Open(filename); err != nil {
		log.Fatalln(err.Error())
	} else {
		defer src.Close()
		buf := make([]byte, 1024)
		if _, err := src.Read(buf); err != nil {
			log.Fatalln(err.Error())
		} else {
			os.Setenv(envName, strings.Trim(string(buf), " \n\t\r\b\x00"))
		}
	}
}

func help() {
	log.Fatalln(`Usage:
gootus <env_dir> <command> ...args`)
}

func execute(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatalln(err.Error())
	} else {
		cmd.Wait()
	}
}
