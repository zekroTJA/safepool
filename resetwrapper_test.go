package safepool

import "testing"

type testObjectNoReset struct {
	Name string
}

func TestWrapper(t *testing.T) {
	sp := New(func() *ResetWrapper[*testObject] {
		return Wrap(&testObject{}, func(v *testObject) {
			v.Name = ""
		})
	})

	o1 := sp.Get()
	o2 := sp.Get()

	o1.Inner.Name = "object1"
	o2.Inner.Name = "object2"

	sp.Put(o1)
	sp.Put(o2)

	o3 := sp.Get()
	o4 := sp.Get()
	o5 := sp.Get()

	if o1 != o3 {
		t.Fatal("new object was created")
	}

	if o2 != o4 {
		t.Fatal("new object was created")
	}

	if o5 == o1 || o5 == o2 {
		t.Fatal("no new object was created")
	}

	if o3.Inner.Name != "" {
		t.Fatal("o1 has not been reset")
	}

	if o4.Inner.Name != "" {
		t.Fatal("o1 has not been reset")
	}
}
