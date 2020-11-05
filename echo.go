package requestid

import "github.com/labstack/echo/v4"

// EchoMiddleware Sets the requestID
func EchoMiddleware() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			requestID := c.Response().Header().Get(headerXRequestID)
			requestIDContext := Set(ctx, requestID)
			c.SetRequest(c.Request().WithContext(requestIDContext))

			return handlerFunc(c)
		}
	}
}
