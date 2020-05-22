package ispbyip

import (
	"sort"
)

func ispv4Lookup(ranges []ISPv4Range, ip uint32) uint32 {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return ranges[i].Last >= ip
	})

	for i := index; i < length; i++ {
		ipRange := ranges[i]

		if ip >= ipRange.First {
			return ipRange.ISP
		}
	}

	return 0
}

func ispv4MultiLookup(ranges []ISPv4Range, ip uint32) []uint32 {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return ranges[i].Last >= ip
	})

	var result []uint32

	for i := index; i < length; i++ {
		ipRange := ranges[i]

		if ip >= ipRange.First {
			result = append(result, ipRange.ISP)
		}
	}

	return result
}
