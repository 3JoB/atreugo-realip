package realip

import (
	"net"
	"net/textproto"
	"strings"

	"github.com/3JoB/atreugo-realip/util"
	"github.com/savsgio/atreugo/v11"
)

var headers = [...]string{
	"X-Client-IP",
	"X-Original-Forwarded-For",
	"X-Forwarded-For",
	"CF-Connecting-IP", // Cloudflare
	"Fastly-Client-Ip", // Fastly CDN and Firebase hosting
	"True-Client-Ip",   // Akamai and Cloudflare
	"X-Real-IP",        // Nginx proxy/FastCGI
	"X-Forwarded",
	"Forwarded-For",
	"Forwarded",
}

func FromRequest(c *atreugo.RequestCtx) string {
	if c == nil {
		return ""
	}

	for _, h := range headers {
		val := util.RequestHeader(c, textproto.CanonicalMIMEHeaderKey(h))

		if strings.ContainsRune(val, ',') {
			for _, address := range strings.Split(val, ",") {
				address = strings.TrimSpace(address)
				if isValidPublicAddress(address) {
					return address
				}
			}
		} else {
			if isValidPublicAddress(val) {
				return val
			}
		}
	}

	remoteAddr := c.RemoteAddr().String()
	var remoteIP string

	if strings.ContainsRune(remoteAddr, ':') {
		remoteIP, _, _ = net.SplitHostPort(remoteAddr)
	} else {
		remoteIP = remoteAddr
	}

	return remoteIP
}

func isValidPublicAddress(addr string) bool {
	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}

	return !IsPrivateIp(ip)
}
