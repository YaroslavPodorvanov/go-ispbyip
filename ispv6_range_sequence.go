package ispbyip

import "sort"

// stay duplicated
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

	return all
}
