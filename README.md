## Max Element

### Get a repo
```bash
>>> go get -u github.com/antik9/gootus@max-element
```

### Usage
```go
package main

import (
	"fmt"
	"github.com/antik9/gootus"
)

type User struct {
	name string
	age  int
}

func main() {
	users := make([]User, 0)
	for i := 1; i <= 10; i++ {
		users = append(users, User{"Mr.Orange", i * 10})
	}
	user, _ := max.FindMax(
		users, func(i, j int) bool { return users[i].age < users[j].age })
	fmt.Println(user)
}
```
