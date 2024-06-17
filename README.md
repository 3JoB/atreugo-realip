Atreugo RealIP
===============
> Realip library adapted for Atreugo.

[![License][license-image]][license-link]
[![Build][build-image]][build-link]

# Example
```go
package main

import (
    "fmt"
    "github.com/savsgio/atreugo/v11"
    "github.com/3JoB/atreugo-realip"
)

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse(fmt.Printf("Your Client IP is: %s", realip.FromRequest(ctx)))
	})
	
	server.GET("/cf", func(ctx *atreugo.RequestCtx) error {
		// realip.County Only available on Cloudflare. 
		// If you need other ways to determine the country, 
		// you need to access the IP database yourself.
		return ctx.TextResponse(fmt.Printf("Your Client IP is: %s, from %s", realip.FromRequest(ctx), realip.County(ctx)))
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
```

[license-link]: https://github.com/3JoB/atreugo-realip/blob/master/LICENSE
[license-image]: https://img.shields.io/dub/l/vibe-d.svg
[build-image]: https://github.com/3JoB/atreugo-realip/actions/workflows/go.yml/badge.svg
[build-link]: https://github.com/3JoB/atreugo-realip/actions