## Parallel Execution

### Get a Repo
```bash
>>> go get -u github.com/antik9/gootus@parallel
```

### Usage

```go
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

    "github.com/antik9/gootus"
)

func printSome() error {
	time.Sleep(time.Second)
	if value := rand.Int(); value%3 == 0 {
		fmt.Println("Error")
		return errors.New("Err")
	} else {
		fmt.Println("OK")
		return nil
	}
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	e, _ := parallel.NewExecutor(5, 3)  // (Number of workers, Maximum errors)
	fs := make([]func() error, 0, 30)

	for i := 0; i < 30; i++ {
		fs = append(fs, printSome)
	}

	e.RunTasks(fs)
}
```
