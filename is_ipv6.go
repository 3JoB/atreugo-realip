package realip

import "net"

func IsIPv6(ip net.IP) bool {
	return IsStringIPv6(ip.String())
}

func IsStringIPv6(ip string) bool {
	for i := 0; i < len(ip); i++ {
		if ':' == ip[i] {
			return true
		}
	}

	return false
}
