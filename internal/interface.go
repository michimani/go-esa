package internal

import (
	"net/http"
)

type IInput interface {
	EsaAPIParameter() (*EsaAPIParameter, error)
}

type IPaginationParameters interface {
	PageValue() (int, bool)
	PerPageValue() (int, bool)
}

type IOutput interface {
	SetRateLimitInfo(h http.Header)
}
