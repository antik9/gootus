## Link Shortener

### Get a repo
```bash
>>> go get -u github.com/antik9/gootus@top-ten
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/antik9/gootus"
)


func main() {
    for _, word := range topten.TopTen("one one two") {
        fmt.Println(word)
    }

    for _, word := range topten.TopN("one one two", 1) {
        fmt.Println(word)
    }
}
```
