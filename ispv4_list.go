package ispbyip

import "sync/atomic"

type ISPv4List struct {
	data atomic.Value
}

func NewISPv4List() *ISPv4List {
	var result = new(ISPv4List)

	result.update(nil)

	return result
}

func (l *ISPv4List) Lookup(ip uint32) uint32 {
	var ranges = l.data.Load().([]ISPv4Range)

	return ispv4Lookup(ranges, ip)
}

func (l *ISPv4List) Update(ranges []ISPv4Range) {
	l.update(ranges)
}

func (l *ISPv4List) update(ranges []ISPv4Range) {
	l.data.Store(ranges)
}
