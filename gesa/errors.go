package gesa

import (
	"errors"
	"fmt"
	"strings"

	"github.com/michimani/go-esa/internal"
)

// GesaError は gesa.Client が返すエラーを表す構造体
type GesaError struct {
	err   error
	OnAPI bool
	EsaAPIError
}

// EsaAPIError は esa API が返すエラーを表す構造体
type EsaAPIError struct {
	Status        string                `json:"-"`
	StatusCode    int                   `json:"-"`
	Error         string                `json:"error"`
	Message       string                `json:"message"`
	RateLimitInfo *RateLimitInformation `json:"-"`
}

func wrapErr(e error) *GesaError {
	if e == nil {
		return nil
	}

	if w, ok := e.(*GesaError); ok {
		return w
	}

	return &GesaError{err: e}
}

func wrapWithAPIErr(eae *EsaAPIError) *GesaError {
	if eae == nil {
		return nil
	}
	return &GesaError{
		err:         errors.New(eae.Summary()),
		OnAPI:       true,
		EsaAPIError: *eae,
	}
}

func (e *GesaError) Error() string {
	if e == nil {
		return ""
	}

	if e.err != nil {
		return e.err.Error()
	}

	return internal.ErrorUndefined
}

func (e *GesaError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

// Summary は esa API が返すエラー内容を要約して string 型で返す
func (e *EsaAPIError) Summary() string {
	if e == nil {
		return ""
	}

	summary := []string{"The esa API returned error response with a status other than 2XX series."}
	if e.Status != "" {
		summary = append(summary, fmt.Sprintf("httpStatus=\"%s\"", e.Status))
	}
	if e.StatusCode > 0 {
		summary = append(summary, fmt.Sprintf("httpStatusCode=%d", e.StatusCode))
	}
	if e.Error != "" {
		summary = append(summary, fmt.Sprintf("error=\"%s\"", e.Error))
	}
	if e.Message != "" {
		summary = append(summary, fmt.Sprintf("message=\"%s\"", e.Message))
	}

	if e.RateLimitInfo != nil {
		summary = append(summary, fmt.Sprintf("rateLimit=%d rateLimitRemaining=%d rateLimitReset=\"%d\"", e.RateLimitInfo.Limit, e.RateLimitInfo.Remaining, e.RateLimitInfo.Reset.SafeTimestamp()))
	}

	return strings.Join(summary, " ")
}
