package plugin

import (
	"context"
	"testing"
)

type testPlugin struct {
	name string
}

// Name return the plugin name
func (t *testPlugin) Name() string {
	return t.name
}

// Perpare called by foundation only once, when the foudation.Perpare func is called
func (t *testPlugin) Perpare(ctx context.Context) error {
	return nil
}

// Run run the plugin anytime which it is needed. However, it may be called many times.
func (t *testPlugin) Run(ctx context.Context) error {
	return nil
}

func Test_GetSet(t *testing.T) {
	f := &foundation{items: make(map[string]int), itemsArr: make([]*item, 0)}
	a := &testPlugin{"a"}
	b := &testPlugin{"b"}

	f.set(a)
	f.set(b)

	if len(f.items) != len(f.itemsArr) && len(f.items) != 2 {
		t.Fatalf("want %v items, actual %v", 2, len(f.items))
	}

	va := f.get("a")

	if va.Name() != a.Name() {
		t.Fatalf("plugin get want %v items, actual %v", a.Name(), va.Name())
	}

	vc := f.get("c")
	if vc != nil {
		t.Fatalf("plugin should be notexsit . want %v items, actual %v", nil, vc)
	}
}
