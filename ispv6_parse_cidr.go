package ispbyip

import (
	"encoding/binary"
	"math"
	"net"
)

func parseIPv6CIDR(cidr string) (IPv6, IPv6, error) {
	_, ipNet, err := net.ParseCIDR(cidr)

	if err != nil {
		return IPv6{}, IPv6{}, err
	}

	if len(ipNet.Mask) == 16 {
		var (
			ip0   = binary.BigEndian.Uint64(ipNet.IP[:8])
			ip1   = binary.BigEndian.Uint64(ipNet.IP[8:])
			mask0 = binary.BigEndian.Uint64(ipNet.Mask[:8])
			mask1 = binary.BigEndian.Uint64(ipNet.Mask[8:])

			first0 = ip0 & mask0
			first1 = ip1 & mask1

			first = IPv6{
				first0,
				first1,
			}

			last = IPv6{
				first0 | (math.MaxUint64 ^ mask0),
				first1 | (math.MaxUint64 ^ mask1),
			}
		)

		return first, last, nil
	}

	return IPv6{}, IPv6{}, ErrExpectIPv6CIDR
}
