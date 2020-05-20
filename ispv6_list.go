package ispbyip

import (
	"sync/atomic"
)

type ISPv6List struct {
	data atomic.Value
}

func NewISPv6List() *ISPv6List {
	var result = new(ISPv6List)

	result.update(nil)

	return result
}

func (l *ISPv6List) Lookup(ip IPv6) uint32 {
	var ranges = l.data.Load().([]ISPv6Range)

	return ispv6Lookup(ranges, ip)
}

func (l *ISPv6List) Update(ranges []ISPv6Range) {
	l.update(ranges)
}

func (l *ISPv6List) update(ranges []ISPv6Range) {
	l.data.Store(ranges)
}
