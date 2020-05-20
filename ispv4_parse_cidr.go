package ispbyip

import (
	"encoding/binary"
	"math"
	"net"
)

func parseIPv4CIDR(cidr string) (uint32, uint32, error) {
	_, ipNet, err := net.ParseCIDR(cidr)

	if err != nil {
		return 0, 0, err
	}

	if len(ipNet.Mask) == 4 {
		var (
			ip    = binary.BigEndian.Uint32(ipNet.IP)
			mask  = binary.BigEndian.Uint32(ipNet.Mask)
			first = ip & mask
			last  = first | (math.MaxUint32 ^ mask)
		)

		return first, last, nil
	}

	return 0, 0, ErrExpectIPv4CIDR
}
