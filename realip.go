package realip

import (
	"net"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
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

func FromRequest(ctx *fasthttp.RequestCtx) string {
	if ctx == nil {
		return ""
	}

	for _, h := range headers {
		val := string(ctx.Request.Header.Peek(http.CanonicalHeaderKey(h)))

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

	remoteAddr := ctx.RemoteAddr().String()
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
