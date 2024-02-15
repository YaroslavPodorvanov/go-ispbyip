# ISP by IP
Lookup ISP by IP

## Terms of use<sup>[[?]](https://github.com/Tyrrrz/.github/blob/master/docs/why-so-political.md)</sup>

By using this project or its source code, for any purpose and in any shape or form, you grant your **implicit agreement** to all the following statements:

- You **condemn Russia and its military aggression against Ukraine**
- You **recognize that Russia is an occupant that unlawfully invaded a sovereign state**
- You **support Ukraine's territorial integrity, including its claims over temporarily occupied territories of Crimea and Donbas**
- You **reject false narratives perpetuated by Russian state propaganda**

To learn more about the war and how you can help, [click here](https://tyrrrz.me/ukraine). Glory to Ukraine! 🇺🇦

### Examples
```golang
package test

import (
	"net"
	"testing"
	
	"github.com/YaroslavPodorvanov/go-ispbyip"

	"github.com/stretchr/testify/require"
)

func TestISPv4List_Lookup(t *testing.T) {
	var ispv4List = ispbyip.NewISPv4List()

	require.Equal(t, uint32(0), ispv4List.Lookup(1))
	require.Equal(t, uint32(0), ispv4List.Lookup(1<<24|1<<16|1<<8+1))
	require.Equal(t, uint32(0), ispv4List.Lookup(2<<24|1<<16|1<<8+2))

	ispv4List.Update(ispbyip.RangeSequenceByISPv4CIDRs([]ispbyip.Input{
		{
			ISP:  1,
			CIDR: "1.1.1.0/24",
		},
		{
			ISP:  2,
			CIDR: "2.1.1.0/24",
		},
	}))

	require.Equal(t, uint32(0), ispv4List.Lookup(1))
	require.Equal(t, uint32(1), ispv4List.Lookup(1<<24|1<<16|1<<8+1))
	require.Equal(t, uint32(2), ispv4List.Lookup(2<<24|1<<16|1<<8+2))
}

func TestList_Lookup(t *testing.T) {
	var (
		ispv4List = ispbyip.NewISPv4List()
		ispv6List = ispbyip.NewISPv6List()
		list      = ispbyip.NewList(ispv4List, ispv6List)
	)

	require.Equal(t, uint32(0), list.Lookup(net.ParseIP("0.0.0.1")))
	require.Equal(t, uint32(0), list.Lookup(net.ParseIP("1:2:3:4::1")))
	require.Equal(t, uint32(0), list.Lookup(net.ParseIP("1:2:3:5::1")))

	ispv4List.Update(ispbyip.RangeSequenceByISPv4CIDRs([]ispbyip.Input{
		{
			ISP:  1,
			CIDR: "0.0.0.1/24",
		},
	}))
	ispv6List.Update(ispbyip.RangeSequenceByISPv6CIDRs([]ispbyip.Input{
		{
			ISP:  2,
			CIDR: "1:2:3:4::1/118",
		},
	}))

	require.Equal(t, uint32(1), list.Lookup(net.ParseIP("0.0.0.1")))
	require.Equal(t, uint32(2), list.Lookup(net.ParseIP("1:2:3:4::1")))
	require.Equal(t, uint32(0), list.Lookup(net.ParseIP("1:2:3:5::1")))
}
```