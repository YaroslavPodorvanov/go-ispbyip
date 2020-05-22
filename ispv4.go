package ispbyip

type ISPv4Range struct {
	ISP uint32
	IPv4Range
}

type IPv4Range struct {
	First uint32
	Last  uint32
}

type ISPv4RangeSort []ISPv4Range

func (c ISPv4RangeSort) Len() int {
	return len(c)
}

func (c ISPv4RangeSort) Less(i, j int) bool {
	a := c[i]
	b := c[j]

	if a.Last < b.Last {
		return true
	}

	if a.Last > b.Last {
		return false
	}

	return a.First < b.First
}

func (c ISPv4RangeSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
