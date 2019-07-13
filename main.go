package main

import (
	"os"
)

func main() {
	if len(os.Args) < 3 {
		help()
	}

	exportEnvFromDir(os.Args[1])
	execute(os.Args[2], os.Args[3:]...)
}
