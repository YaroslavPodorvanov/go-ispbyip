package ispbyip

import "sort"

func RangeSequenceByISPv6CIDRs(ispcidrs []Input) []ISPv6Range {
	length := len(ispcidrs)

	if length == 0 {
		return nil
	}

	all := make([]ISPv6Range, 0, length)

	for _, ispcidr := range ispcidrs {
		first, last, err := parseIPv6CIDR(ispcidr.CIDR)

		if err != nil {
			// NOP

			continue
		}

		all = append(all, ISPv6Range{
			ISP: ispcidr.ISP,
			IPv6Range: IPv6Range{
				First: first,
				Last:  last,
			},
		})
	}

	// O(N*ln(N))
	sort.Sort(ISPv6RangeSort(all))

	unique := all[:1]

	current := all[0]

	// O(N)
	for i := 1; i < length; i++ {
		next := all[i]

		if IPv6Less(current.Last, next.Last) {
			unique = append(unique, next)

			current = next
		}
	}

	return unique
}
