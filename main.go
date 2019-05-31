package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err == nil {
		fmt.Println(time.String())
	} else {
		fmt.Println(err)
	}
}
