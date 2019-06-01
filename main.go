package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err == nil {
		fmt.Println(time.String())
	} else {
		log.Println(err)
		os.Exit(1)
	}
}
