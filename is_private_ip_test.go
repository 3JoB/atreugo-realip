package realip

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrivateIpV4(t *testing.T) {
	cases := []struct {
		ip        net.IP
		isPrivate bool
		msg       string
	}{
		{ip: net.ParseIP("127.0.0.1"), isPrivate: true, msg: "127.0.0.1 should be private"},
		{ip: net.ParseIP("192.168.254.254"), isPrivate: true, msg: "192.168.254.254 should be private"},
		{ip: net.ParseIP("10.255.0.3"), isPrivate: true, msg: "10.255.0.3 should be private"},
		{ip: net.ParseIP("172.16.255.255"), isPrivate: true, msg: "172.16.255.255 should be private"},
		{ip: net.ParseIP("172.31.255.255"), isPrivate: true, msg: "172.31.255.255 should be private"},
		{ip: net.ParseIP("192.169.255.255"), isPrivate: false, msg: "192.169.255.255 should not be private"},
		{ip: net.ParseIP("9.255.0.255"), isPrivate: false, msg: "9.255.0.255 should not be private"},
		{ip: net.ParseIP("67.22.74.181"), isPrivate: false, msg: "67.22.74.181 should not be private"},
		{ip: net.ParseIP("109.86.51.55"), isPrivate: false, msg: "109.86.51.55 should not be private"},
		{ip: net.ParseIP("76.233.86.21"), isPrivate: false, msg: "76.233.86.21 should not be private"},
		{ip: net.ParseIP("109.42.112.5"), isPrivate: false, msg: "109.42.112.5 should not be private"},
		{ip: net.ParseIP("81.185.175.46"), isPrivate: false, msg: "81.185.175.46 should not be private"},
	}

	for _, ts := range cases {
		assert.Equal(t, ts.isPrivate, IsPrivateIp(ts.ip))
	}
}

func TestIsPrivateIpV6(t *testing.T) {
	cases := []struct {
		ip        net.IP
		isPrivate bool
		msg       string
	}{
		{ip: net.ParseIP("::0"), isPrivate: true, msg: "::0 should be private"},
		{ip: net.ParseIP("::1"), isPrivate: true, msg: "::1 should be private"},
		{ip: net.ParseIP("fe80::1"), isPrivate: true, msg: "fe80::1 should be private"},
		{ip: net.ParseIP("febf::1"), isPrivate: true, msg: "febf::1 should be private"},
		{ip: net.ParseIP("ff00::1"), isPrivate: true, msg: "ff00::1 should be private"},
		{ip: net.ParseIP("ff10::1"), isPrivate: true, msg: "ff10::1 should be private"},
		{ip: net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"), isPrivate: true, msg: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff should be private"},
		{ip: net.ParseIP("2002::"), isPrivate: true, msg: "2002:: should be private"},
		{ip: net.ParseIP("2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff"), isPrivate: true, msg: "2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff should be private"},

		{ip: net.ParseIP("::2"), isPrivate: false, msg: "::2 should not be private"},
		{ip: net.ParseIP("fec0::1"), isPrivate: false, msg: "fec0::1 should not be private"},
		{ip: net.ParseIP("feff::1"), isPrivate: false, msg: "feff::1 should not be private"},
		{ip: net.ParseIP("2409:4073:f:11b9:817a:81ea:a128:1955"), isPrivate: false, msg: "2409:4073:f:11b9:817a:81ea:a128:1955 should not be private"},
		{ip: net.ParseIP("2401:4900:1a8d:3eaf:84b5:5538:efd:5586"), isPrivate: false, msg: "2401:4900:1a8d:3eaf:84b5:5538:efd:5586 should not be private"},
		{ip: net.ParseIP("2409:4063:230c:e172:dd1a:4c74:de7c:cc21"), isPrivate: false, msg: "2409:4063:230c:e172:dd1a:4c74:de7c:cc21 should not be private"},
		{ip: net.ParseIP("2601:240:a:1bc6:9c39:66f0:ced8:d643"), isPrivate: false, msg: "2601:240:a:1bc6:9c39:66f0:ced8:d643 should not be private"},
	}

	for _, ts := range cases {
		assert.Equal(t, ts.isPrivate, IsPrivateIp(ts.ip))
	}
}
