package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Follow link to see details
// https://leetcode.com/problems/validate-ip-address/

const (
	Neither = "Neither"
	IPv4    = "IPv4"
	IPv6    = "IPv6"

	maxNum = 1 << 16
)

func validIPAddress(queryIP string) string {
	if strings.Contains(queryIP, ".") {
		return validIPv4Address(queryIP)
	}

	return validIPv6Address(queryIP)
}

func validIPv4Address(queryIP string) string {
	parts := strings.Split(queryIP, ".")
	if len(parts) != 4 {
		return Neither
	}

	for _, part := range parts {
		if len(part) == 0 || len(part) > 3 || (part[0] == '0' && len(part) > 1) {
			return Neither
		}

		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return Neither
		}
	}

	return IPv4
}

func validIPv6Address(queryIP string) string {
	parts := strings.Split(queryIP, ":")
	if len(parts) != 8 {
		return Neither
	}

	for _, part := range parts {
		if len(part) == 0 || len(part) > 4 {
			return Neither
		}

		v, err := strconv.ParseInt(part, 16, 64)
		if err != nil || v < 0 || v > maxNum {
			return Neither
		}
	}

	return IPv6
}

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))                      // IPv4
	fmt.Println(validIPAddress("01.01.01.01"))                       // Neither
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334")) // IPv6
	fmt.Println(validIPAddress("256.256.256.256"))                   // Neither
}
