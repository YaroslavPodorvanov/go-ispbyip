package ispbyip

type IPv6 [2]uint64

type ISPv6Range struct {
	ISP uint32
	IPv6Range
}

type IPv6Range struct {
	First IPv6
	Last  IPv6
}

type ISPv6RangeSort []ISPv6Range

func (c ISPv6RangeSort) Len() int {
	return len(c)
}

func (c ISPv6RangeSort) Less(i, j int) bool {
	a := c[i]
	b := c[j]

	if IPv6Less(a.First, b.First) {
		return true
	}

	if IPv6Less(b.First, a.First) {
		return false
	}

	// 0 to 127, before 0 to 15
	return IPv6Less(b.Last, a.Last)
}

func (c ISPv6RangeSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func IPv6Less(a, b IPv6) bool {
	if a[0] > b[0] {
		return false
	}

	return a[1] < b[1]
}

func IPv6LessEqual(a, b IPv6) bool {
	if a[0] > b[0] {
		return false
	}

	return a[1] <= b[1]
}
