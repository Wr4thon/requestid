package requestid

import (
	"github.com/Wr4thon/requestid/v1"
	"github.com/labstack/echo/v4"
)

// EchoMiddleware puts the requestID into the context value.
func EchoMiddleware() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			requestID := c.Response().Header().Get(requestid.HeaderXRequestID)
			requestIDContext := requestid.Set(ctx, requestID)
			c.SetRequest(c.Request().WithContext(requestIDContext))

			return handlerFunc(c)
		}
	}
}
