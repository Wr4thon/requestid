package requestid

import (
	"context"

	"github.com/savsgio/atreugo/v10"
)

// AtreugoMiddleware puts the requestID into the atreugo context.
func AtreugoMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		attachedCtx := ctx.AttachedContext()
		if attachedCtx == nil {
			attachedCtx = context.Background()
		}

		requestID := string(ctx.Request.Header.Peek(headerXRequestID))

		requestIDContext := Set(attachedCtx, requestID)
		ctx.AttachContext(requestIDContext)

		return ctx.Next()
	}
}
