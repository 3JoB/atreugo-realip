package realip

import (
	"net"
	"net/textproto"
	"strings"

	"github.com/3JoB/atreugo-realip/util"
	"github.com/savsgio/atreugo/v11"
)

var headers = [...]string{
	"CF-Connecting-IP", // Cloudflare
	"Fastly-Client-Ip", // Fastly CDN and Firebase hosting
	"True-Client-Ip",   // Akamai and Cloudflare
	"X-Real-IP",        // Nginx proxy/FastCGI
	"X-Client-IP",
	"X-Original-Forwarded-For",
	"X-Forwarded-For",
	"X-Forwarded",
	"Forwarded-For",
	"Forwarded",
}

// Only works when using Cloudflare as CDN!
func County(c *atreugo.RequestCtx) string {
	if c == nil {
		return ""
	}
	return util.RequestHeader(c, textproto.CanonicalMIMEHeaderKey("Cf-Ipcountry"))
}

func FromRequest(c *atreugo.RequestCtx) string {
	if c == nil {
		return ""
	}

	for _, h := range headers {
		val := util.RequestHeader(c, textproto.CanonicalMIMEHeaderKey(h))
		if strings.ContainsRune(val, ',') {
			str_sp := strings.Split(val, ",")
			str_len := len(str_sp)
			address := ""
			if str_len > 2 {
				address = strings.TrimSpace(str_sp[1])
				
			} else {
				address = strings.TrimSpace(str_sp[0])
			}
			if isValidPublicAddress(address) {
				return address
			}
		} else {
			if isValidPublicAddress(val) {
				return val
			}
		}
	}

	remoteAddr := util.RequestHeader(c, textproto.CanonicalMIMEHeaderKey("Remote-Host"))
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
