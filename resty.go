package requestid

import (
	resty "github.com/go-resty/resty/v2"
)

const (
	headerXRequestID = "X-Request-ID"
)

// RestyMiddleware gets the requestID from the context
// and puts it into the header.
func RestyMiddleware(c *resty.Client, r *resty.Request) error {
	rID, err := Get(r.Context())
	if err == nil {
		r.SetHeader(headerXRequestID, rID)
	}

	return nil
}
