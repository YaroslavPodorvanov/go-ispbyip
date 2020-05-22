package ispbyip

import "sort"

// stay duplicated
func RangeSequenceByISPv4CIDRs(ispcidrs []Input) []ISPv4Range {
	length := len(ispcidrs)

	if length == 0 {
		return nil
	}

	all := make([]ISPv4Range, 0, length)

	for _, ispcidr := range ispcidrs {
		first, last, err := parseIPv4CIDR(ispcidr.CIDR)

		if err != nil {
			// NOP

			continue
		}

		all = append(all, ISPv4Range{
			ISP: ispcidr.ISP,
			IPv4Range: IPv4Range{
				First: first,
				Last:  last,
			},
		})
	}

	// O(N*ln(N))
	sort.Sort(ISPv4RangeSort(all))

	return all
}
