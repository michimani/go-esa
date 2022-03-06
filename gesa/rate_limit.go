package gesa

import (
	"net/http"
	"strconv"

	"github.com/michimani/go-esa/internal"
)

const (
	RATE_LIMIT_LIMIT_HEADER_KEY     = "X-RateLimit-Limit"
	RATE_LIMIT_REMAINING_HEADER_KEY = "X-RateLimit-Remaining"
	RATE_LIMIT_RESET_HEADER_KEY     = "X-RateLimit-Reset"
)

func GetRateLimitInformation(resHeader http.Header) (*RateLimitInformation, error) {
	i := RateLimitInformation{}
	limit, err := rateLimitLimit(resHeader)
	if err != nil {
		return nil, err
	}
	i.Limit = limit

	remaining, err := rateLimitRemaining(resHeader)
	if err != nil {
		return nil, err
	}
	i.Remaining = remaining

	reset, err := rateLimitReset(resHeader)
	if err != nil {
		return nil, err
	}
	i.Reset = reset

	return &i, nil
}

func rateLimitLimit(h http.Header) (int, error) {
	values := internal.HeaderValues(RATE_LIMIT_LIMIT_HEADER_KEY, h)
	if len(values) == 0 {
		return 0, nil
	}
	limit, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, err
	}

	return limit, nil
}

func rateLimitRemaining(h http.Header) (int, error) {
	values := internal.HeaderValues(RATE_LIMIT_REMAINING_HEADER_KEY, h)
	if len(values) == 0 {
		return 0, nil
	}
	remaining, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, err
	}

	return remaining, nil
}

func rateLimitReset(h http.Header) (*Timestamp, error) {
	values := internal.HeaderValues(RATE_LIMIT_RESET_HEADER_KEY, h)
	if len(values) == 0 {
		return nil, nil
	}
	resetInt, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, err
	}

	ts := Timestamp(resetInt)
	return &ts, nil
}
