# requestid

## Description

this repository is a collection of middlewares, used to get a requestID from an http request, and set it in rest clients used to perform requests.

## supported frameworks

| Framework | Support | Version |
|:----------|:-------:|:--------:|
| Resty     |    ✓    |      v2 |
| Echo      |    ✓    |      v4 |
| Atreugo   |    ✓    |     v10 |

## how to install

```bash
go get github.com/Wr4thon/requestid@v0.1.0
```

## How to use

### resty

```go
package main

import (
	"context"

	"github.com/Wr4thon/requestid"
	"github.com/go-resty/resty/v2"
)

func main(){
	client := resty.
		New().
		OnBeforeRequest(requestid.RestyMiddleware)

	// this most likely will be set by an http framework such as atreugo or echo.
	ctx := context.TODO()
	ctx = requestid.Set(ctx, "<TheRequestID>")

	resp, err := client.R().
		SetContext(ctx).
		Get("https://example.com")

	if err != nil {
		panic(err)
	}

	println(string(resp.Body()))
}

```

### echo

```go
package main

import (
	"net/http"

	"github.com/Wr4thon/requestid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(requestid.EchoMiddleware())

	e.GET("/", func(c echo.Context) error {
		rid, err := requestid.Get(c.Request().Context())

		if err != nil {
      println(err.Error())
			return c.NoContent(http.StatusInternalServerError)
		}

		println(rid)
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Error(e.Start(":1337"))
}

```
