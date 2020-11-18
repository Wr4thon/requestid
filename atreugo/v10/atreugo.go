package requestid

import (
	"context"

	"github.com/Wr4thon/requestid/v1"
	"github.com/savsgio/atreugo/v10"
)

// AtreugoMiddleware puts the requestID into the context value.
func AtreugoMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		attachedCtx := ctx.AttachedContext()
		if attachedCtx == nil {
			attachedCtx = context.Background()
		}

		requestID := string(ctx.Request.Header.Peek(requestid.HeaderXRequestID))

		requestIDContext := requestid.Set(attachedCtx, requestID)
		ctx.AttachContext(requestIDContext)

		return ctx.Next()
	}
}
