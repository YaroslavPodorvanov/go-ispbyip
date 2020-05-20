package ispbyip

import (
	"encoding/binary"
	"net"
)

type List struct {
	ispv4 *ISPv4List
	ispv6 *ISPv6List
}

func NewList(ispv4 *ISPv4List, ispv6 *ISPv6List) *List {
	return &List{ispv4: ispv4, ispv6: ispv6}
}

func (l *List) Lookup(ip net.IP) uint32 {
	{
		var ISPv4 = ip.To4()

		if ISPv4 != nil {
			return l.ispv4.Lookup(binary.BigEndian.Uint32(ISPv4))
		}
	}

	// just in case condition
	{
		var ISPv6 = ip.To16()

		if ISPv6 != nil {
			return l.ispv6.Lookup(IPv6{
				binary.BigEndian.Uint64(ISPv6[:8]),
				binary.BigEndian.Uint64(ISPv6[8:]),
			})
		}
	}

	return 0
}
