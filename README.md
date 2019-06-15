## Double Linked List

### Get a Repo

```bash
>>> go get -u github.com/antik9/gootus@double-linked-list
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/antik9/gootus"
)

func main() {
	newList := list.NewList()

	newList.PushBack(10)
	newList.PushFront(5)
	newList.PushFront(4)

	newList.First().Remove()

	fmt.Println(newList.Len())
	fmt.Println(newList.First().Value())
	fmt.Println(newList.Last().Value())
}
```
