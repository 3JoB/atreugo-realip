FastHTTP RealIP
===============
> Go package that can be used to get client's real public IP from Fast HTTP request.

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