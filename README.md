FastHTTP RealIP
===============
> Go package that can be used to get client's real public IP from Fast HTTP request.

[![License][license-image]][license-link]
[![Build][build-image]][build-link]

# Example
```go
package main

import (
    "log"
    "github.com/valyala/fasthttp"
    "github.com/zhooravell/fasthttp-realip"
)

func main() {
	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
    log.Println("Client IP: " + realip.FromRequest(ctx))
}
```

[license-link]: https://github.com/zhooravell/fasthttp-realip/blob/master/LICENSE
[license-image]: https://img.shields.io/dub/l/vibe-d.svg
[build-image]: https://github.com/zhooravell/fasthttp-realip/actions/workflows/go.yml/badge.svg
[build-link]: https://github.com/zhooravell/fasthttp-realip/actions