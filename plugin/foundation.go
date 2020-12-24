package plugin

import (
	"context"
	"fmt"
	"sync"
)

var (
	def *foundation
)

func init() {
	def = &foundation{items: make(map[string]int), itemsArr: make([]*item, 0)}
	fmt.Println("Wow!!!plugin foudation inited")
}

const (
	PluginNotRegistered = ErrPlugin("the plugin is not registered")
)

// ErrPlugin plugin error type
type ErrPlugin string

// Error imp error
func (pe ErrPlugin) Error() string {
	return string(pe)
}

// Pluggable Pluggable interface
type Pluggable interface {
	// Name return the plugin name
	Name() string
	// Perpare called by foundation only once, when the foudation.Perpare func is called
	Perpare(ctx context.Context) error
	// Run run the plugin anytime which it is needed. However, it may be called many times.
	Run(ctx context.Context) error
}

type item struct {
	Plugin Pluggable
}

type foundation struct {
	items    map[string]int
	itemsArr []*item
	perpare  sync.Once
}

// RunPluginByName call a plugin Run with the plugin name
func (f *foundation) RunPluginByName(ctx context.Context, name string) error {
	it := f.get(name)
	if it == nil {
		return PluginNotRegistered
	}

	return it.Run(ctx)
}

// Run run all plugins
func (f *foundation) Run(ctx context.Context, names ...string) error {
	if len(names) > 0 {
		for _, name := range names {
			err := f.RunPluginByName(ctx, name)

			if err != nil {
				return err
			}
		}
	}

	for _, it := range f.itemsArr {
		err := it.Plugin.Run(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Perpare call all plugins perpare func
func (f *foundation) Perpare(ctx context.Context) error {
	var res error

	do := func() {
		for _, it := range f.itemsArr {
			err := it.Plugin.Perpare(ctx)

			if err != nil {
				res = err
				break
			}
		}
	}

	f.perpare.Do(do)

	return res
}

// Register register plugins
func (f *foundation) Register(ps ...Pluggable) {
	for _, p := range ps {
		f.set(p)
	}
}

func (f *foundation) set(p Pluggable) {
	var it *item

	idx, ok := f.items[p.Name()]

	if !ok {
		//TODO using pool here ?
		it = &item{}
		idx = len(f.itemsArr)
		f.itemsArr = append(f.itemsArr, it)
		f.items[p.Name()] = idx
	} else {
		it = f.itemsArr[idx]
	}

	it.Plugin = p
}

func (f *foundation) get(name string) Pluggable {
	idx, ok := f.items[name]

	if !ok {
		return nil
	}

	it := f.itemsArr[idx]

	return it.Plugin
}
