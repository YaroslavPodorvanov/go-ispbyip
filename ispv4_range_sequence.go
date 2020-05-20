package ispbyip

import "sort"

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

	unique := all[:1]

	current := all[0]

	// O(N)
	for i := 1; i < length; i++ {
		next := all[i]

		if next.Last > current.Last {
			unique = append(unique, next)

			current = next
		}
	}

	return unique
}
