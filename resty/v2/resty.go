package requestid

import (
	"github.com/Wr4thon/requestid/v1"
	resty "github.com/go-resty/resty/v2"
)

// RestyMiddleware gets the requestID from the context
// and puts it into the header.
func RestyMiddleware(c *resty.Client, r *resty.Request) error {
	rID, err := Get(r.Context())
	if err == nil {
		r.SetHeader(requestid.HeaderXRequestID, rID)
	}

	return nil
}
