package internal

import (
	"io"
	"net/http"
)

type IParameters interface {
	Body() (io.Reader, error)
	ResolveEndpoint(endpointBase string) string
	ParameterMap() map[string]string
}

type IPaginationParameters interface {
	PageValue() (int, bool)
	PerPageValue() (int, bool)
}

type IResponse interface {
	SetRateLimitInfo(h http.Header)
}
