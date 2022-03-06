package internal

import (
	"io"
	"net/http"
)

type IParameters interface {
	Body() (io.Reader, error)
}

type IResponse interface {
	SetRateLimitInfo(h http.Header)
}
