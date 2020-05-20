package ispbyip

import "sort"

func ispv6Lookup(ranges []ISPv6Range, ip IPv6) uint32 {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return IPv6LessEqual(ip, ranges[i].Last)
	})

	if index < length {
		ipRange := ranges[index]

		if IPv6LessEqual(ipRange.First, ip) {
			return ipRange.ISP
		}
	}

	return 0
}
