## Homework Events

### Get a Repo

```bash
>>> go get -u github.com/antik9/gootus@log-events
```

### Usage

```go
package main

import (
    "os"
	"github.com/antik9/gootus"
)

func main() {
	submitted := events.HwSubmitted{1, "func main() { }", "My go function"}
	events.LogOtusEvent(submitted, os.Stdout)

	accepted := events.HwAccepted{1, 5}
	events.LogOtusEvent(accepted, os.Stdout)
}
```
