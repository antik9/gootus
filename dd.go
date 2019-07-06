package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	// DefaultOffset is used if no offset is given from command line
	DefaultOffset = 0
	// DefaultLimit is used if no limit is given from command line
	DefaultLimit = math.MaxInt64
	batchSize    = int64(1024 * 1024)
)

// Copy from src to dest from offset limit bytes
func Copy(src, dest string, offset, limit int64) {
	fmt.Println(src, dest, offset, limit)
	sizeOfSrc := sizeOf(src)

	if sizeOfSrc < offset {
		log.Fatalln("offset is too big for this file")
	}
	if limit > sizeOfSrc || offset+limit > sizeOfSrc {
		limit = sizeOfSrc - offset
	}
	log.Printf("will copy %d bytes\n", limit)

	makeCopy(src, dest, offset, limit)
}

func failOnErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func makeCopy(src, dest string, offset, limit int64) {
	var err error
	initialSize := limit

	srcFile := openFile(src)
	defer srcFile.Close()

	srcFile.Seek(offset, io.SeekStart)

	destFile := openOrCreateFile(dest)
	defer destFile.Close()

	buf := make([]byte, batchSize)

	for limit > 0 {
		nextBatchSize := batchSize

		if limit < batchSize {
			nextBatchSize = limit
		}

		_, err = srcFile.Read(buf)
		failOnErr(err)

		_, err = destFile.Write(buf[:nextBatchSize])
		failOnErr(err)

		limit -= nextBatchSize
		printStats(initialSize, initialSize-limit)
	}
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	failOnErr(err)
	return file
}

func openOrCreateFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	failOnErr(err)
	return file
}

func printStats(initialSize, alreadyCopied int64) {
	fmt.Printf("\b\b\b\b\b\b\b%.2f%%", float64(alreadyCopied)/float64(initialSize)*100.0)
}

func sizeOf(filename string) int64 {
	file := openFile(filename)
	defer file.Close()

	stat, err := file.Stat()
	failOnErr(err)

	return stat.Size()
}
