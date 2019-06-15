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
)

func main() {
	list := NewList()

	list.PushBack(10)
	list.PushFront(5)
	list.PushFront(4)

	list.First().Remove()

	fmt.Println(list.Len())
	fmt.Println(list.First().Value())
	fmt.Println(list.Last().Value())
}
```
