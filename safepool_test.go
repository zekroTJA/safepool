package safepool

import "testing"

type testObject struct {
	Name string
}

func (t *testObject) ResetState() {
	t.Name = ""
}

func Test(t *testing.T) {
	sp := New(func() *testObject {
		return &testObject{}
	})

	o1 := sp.Get()
	o2 := sp.Get()

	o1.Name = "object1"
	o2.Name = "object2"

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

	if o3.Name != "" {
		t.Fatal("o1 has not been reset")
	}

	if o4.Name != "" {
		t.Fatal("o1 has not been reset")
	}
}
