package util

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"
)

func RequestHeader(c *atreugo.RequestCtx, key string) string {
	return unsafeConvert.StringSlice(c.Request.Header.Peek(key))
}

func ResponseHeader(c *atreugo.RequestCtx, key string) string {
	return unsafeConvert.StringSlice(c.Response.Header.Peek(key))
}
