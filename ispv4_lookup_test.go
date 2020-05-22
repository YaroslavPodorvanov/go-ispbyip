package ispbyip

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestISPv4Lookup(t *testing.T) {
	require.Equal(t, uint32(0), ispv4Lookup(nil, 1))
	require.Equal(t, uint32(0), ispv4Lookup(nil, 2))
	require.Equal(t, uint32(0), ispv4Lookup(nil, 3))

	var ranges = []ISPv4Range{
		{
			ISP: 1,
			IPv4Range: IPv4Range{
				First: 0,
				Last:  127,
			},
		},
		{
			ISP: 2,
			IPv4Range: IPv4Range{
				First: 128,
				Last:  255,
			},
		},
	}

	assertSort(t, ranges)

	require.Equal(t, uint32(1), ispv4Lookup(ranges, 1))
	require.Equal(t, uint32(1), ispv4Lookup(ranges, 2))
	require.Equal(t, uint32(1), ispv4Lookup(ranges, 3))

	require.Equal(t, uint32(2), ispv4Lookup(ranges, 128))
	require.Equal(t, uint32(2), ispv4Lookup(ranges, 192))
	require.Equal(t, uint32(2), ispv4Lookup(ranges, 255))

	require.Equal(t, uint32(0), ispv4Lookup(ranges, 256))
}

func TestISPv4MultiLookup(t *testing.T) {
	require.Equal(t, []uint32(nil), ispv4MultiLookup(nil, 1))
	require.Equal(t, []uint32(nil), ispv4MultiLookup(nil, 2))
	require.Equal(t, []uint32(nil), ispv4MultiLookup(nil, 3))

	var ranges = []ISPv4Range{
		{
			ISP: 1,
			IPv4Range: IPv4Range{
				First: 0,
				Last:  127,
			},
		},
		{
			ISP: 4,
			IPv4Range: IPv4Range{
				First: 1152,
				Last:  1279,
			},
		},
		{
			ISP: 5,
			IPv4Range: IPv4Range{
				First: 1280,
				Last:  1407,
			},
		},
		{
			ISP: 2,
			IPv4Range: IPv4Range{
				First: 1024,
				Last:  1535,
			},
		},
		{
			ISP: 6,
			IPv4Range: IPv4Range{
				First: 1408,
				Last:  1535,
			},
		},
		{
			ISP: 3,
			IPv4Range: IPv4Range{
				First: 1024,
				Last:  2047,
			},
		},
	}

	assertSort(t, ranges)

	require.Equal(t, []uint32{1}, ispv4MultiLookup(ranges, 1))
	require.Equal(t, []uint32{1}, ispv4MultiLookup(ranges, 2))
	require.Equal(t, []uint32{1}, ispv4MultiLookup(ranges, 3))

	require.Equal(t, []uint32{2, 3}, ispv4MultiLookup(ranges, 1024))
	require.Equal(t, []uint32{4, 2, 3}, ispv4MultiLookup(ranges, 1152))

	require.Equal(t, []uint32(nil), ispv4MultiLookup(ranges, 256))
}

func assertSort(t *testing.T, source []ISPv4Range) {
	t.Helper()

	var sorted = make([]ISPv4Range, len(source))
	copy(sorted, source)

	sort.Sort(ISPv4RangeSort(sorted))

	require.Equal(t, source, sorted)
}
