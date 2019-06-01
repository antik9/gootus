## Link Shortener

### Get a repo
```bash
>>> go get -u github.com/antik9/gootus:shortener
```

### Usage

```go
import (
    "fmt"
    "shortener"
)


func main() {
    worker := shortner.LinkShortener{map[string]string{}, map[string]string{}}
    fmt.Println(worker.Shorten("https://google.com")
}
```
