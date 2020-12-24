package plugin

import "sync/atomic"

// stats elem
const (
	IDLE uint32 = iota + 1
	PERPARED
	RUNNING
	ABORT
)

// Stats stats type
type Stats struct {
	flag uint32
}

// Set set stats
func (s *Stats) Set(stats uint32) {
	atomic.StoreUint32(&s.flag, stats)
}

// Equal equal target stats
func (s *Stats) Equal(stats uint32) bool {
	tmp := atomic.LoadUint32(&s.flag)

	return tmp == stats
}

// Swap swap target if stats equal to old
func (s *Stats) Swap(old, target uint32) bool {
	return atomic.CompareAndSwapUint32(&s.flag, old, target)
}
