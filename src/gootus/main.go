package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"topten"
	"unpacker"
)

func help() {
	fmt.Println(`Provide one of the commands:
	gootus unpack # Unpacks string like u3i9
	gootus top 5  # find top 5 occurences from the provided line`)
	os.Exit(1)
}

func unpack() {
	var input string
	fmt.Scan(&input)
	fmt.Println(unpacker.Unpack(input))
}

func printTopN(n int) {
	reader := bufio.NewReader(os.Stdin)
	if input, err := reader.ReadString('\n'); err != nil {
		log.Fatalln(err)
	} else {
		for _, word := range topten.TopN(input, n) {
			fmt.Println(word)
		}
	}
}

func main() {
	if len(os.Args) < 2 || (os.Args[1] != "unpack" && os.Args[1] != "top") {
		help()
	}

	if os.Args[1] == "unpack" {
		unpack()
	} else {
		if len(os.Args) != 3 {
			help()
		}
		if n, err := strconv.Atoi(os.Args[2]); err != nil || n < 0 {
			help()
		} else {
			printTopN(n)
		}
	}
}
