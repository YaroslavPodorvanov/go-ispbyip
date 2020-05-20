package ispbyip

import "errors"

var (
	ErrExpectIPv4CIDR = errors.New("expect IPv4 CIDR")
	ErrExpectIPv6CIDR = errors.New("expect IPv6 CIDR")
)
