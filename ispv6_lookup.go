package ispbyip

import "sort"

func ispv6Lookup(ranges []ISPv6Range, ip IPv6) uint32 {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return IPv6LessEqual(ip, ranges[i].Last)
	})

	for i := index; i < length; i++ {
		ipRange := ranges[i]

		if IPv6LessEqual(ipRange.First, ip) {
			return ipRange.ISP
		}
	}

	return 0
}

func ispv6MultiLookup(ranges []ISPv6Range, ip IPv6) []uint32 {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return IPv6LessEqual(ip, ranges[i].Last)
	})

	var result []uint32

	for i := index; i < length; i++ {
		ipRange := ranges[i]

		if IPv6LessEqual(ipRange.First, ip) {
			result = append(result, ipRange.ISP)
		}
	}

	return result
}
