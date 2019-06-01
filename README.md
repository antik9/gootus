## Link Shortener

### Get a repo
```bash
>>> go get -u github.com/antik9/gootus@shortener
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/antik9/gootus"
)


func main() {
    worker := shortener.NewLinkShortener()
    fmt.Println(worker.Shorten("https://google.com"))
}
```
