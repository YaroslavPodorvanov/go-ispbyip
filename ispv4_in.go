package ispbyip

import "sort"

func ispv4Lookup(ranges []ISPv4Range, ip uint32) uint32 {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return ip <= ranges[i].Last
	})

	if index < length {
		ipRange := ranges[index]

		if ipRange.First <= ip {
			return ipRange.ISP
		}
	}

	return 0
}
