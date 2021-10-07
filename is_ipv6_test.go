package realip

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIPv6(t *testing.T) {
	cases := []struct {
		ip     net.IP
		isIpv6 bool
	}{
		{ip: net.ParseIP("99.250.150.130"), isIpv6: false},
		{ip: net.ParseIP("12.216.74.218"), isIpv6: false},
		{ip: net.ParseIP("76.217.50.159"), isIpv6: false},
		{ip: net.ParseIP("67.22.74.181"), isIpv6: false},
		{ip: net.ParseIP("2409:4055:2e17:5d60:e0bc:c4cc:a9de:fc63"), isIpv6: true},
		{ip: net.ParseIP("2a04:cec0:11fa:45db:7850:ec91:90a5:68a9"), isIpv6: true},
		{ip: net.ParseIP("2a02:1810:4821:cb00:4482:87c5:e8bd:47a2"), isIpv6: true},
		{ip: net.ParseIP("2401:4900:47d9:f521:ed48:4148:22ee:2966"), isIpv6: true},
	}

	for _, testCase := range cases {
		assert.Equal(t, testCase.isIpv6, IsIPv6(testCase.ip))
	}
}
