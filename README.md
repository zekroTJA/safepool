# safepool

This package provides a wrapper around [sync.Pool](https://pkg.go.dev/sync#Pool) which 
ensures type and state safety.

Because the type constraint on the pool, you can always rely to retrieve the expected
type specified to the pool. Also, all objects put into the pool need to implement
ResetState which sets the state of the objects to a clean "zero" state on putting it
back into the pool.

## Example

```go
package main

import (
	"fmt"

	"github.com/zekrotja/safepool"
)

type User struct {
	Name string
}

func (t *User) ResetState() {
	t.Name = ""
}

func main() {
	sp := safepool.New(func() *User {
		return &User{}
	})

	u1 := sp.Get()
	defer sp.Put(u1)

	u2 := sp.Get()
	defer sp.Put(u2)

	u1.Name = "Zero Two"
	u2.Name = "Ichigo"

    // ...
}
```