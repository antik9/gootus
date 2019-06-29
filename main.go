package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"topten"
	"unpacker"
)

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
	runUnpack := flag.Bool("unpack", false, "gootus -unpack # Unpacks string like u3i9")
	topN := flag.Int("top", 0, "gootus -top 5  # find top 5 occurences from the provided line")
	flag.Parse()

	if *runUnpack {
		unpack()
	} else if *topN > 0 {
		printTopN(*topN)
	} else {
		flag.PrintDefaults()
	}
}
