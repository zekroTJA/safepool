package safepool_test

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

func Example() {
	sp := safepool.New(func() *User {
		return &User{}
	})

	u1 := sp.Get()
	defer sp.Put(u1)

	u2 := sp.Get()
	defer sp.Put(u2)

	u1.Name = "Zero Two"
	u2.Name = "Ichigo"

	fmt.Printf("u1: %+v\n", u1)
	fmt.Printf("u2: %+v\n", u2)

	// Output:
	// u1: &{Name:Zero Two}
	// u2: &{Name:Ichigo}
}
