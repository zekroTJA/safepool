package safepool_test

import (
	"fmt"

	"github.com/zekrotja/safepool"
)

type ExternalUser struct {
	Name string
}

func ExampleResetWrapper() {
	sp := safepool.New(func() *safepool.ResetWrapper[*ExternalUser] {
		return safepool.Wrap(&ExternalUser{}, func(v *ExternalUser) {
			v.Name = ""
		})
	})

	user := sp.Get()
	defer sp.Put(user)

	user.Inner.Name = "Zero Two"

	fmt.Printf("user.Inner: %+v\n", user.Inner)

	// Output:
	// user.Inner: &{Name:Zero Two}
}
