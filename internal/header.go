package internal

import (
	"net/http"
	"strings"
)

func HeaderValues(key string, h http.Header) []string {
	if hv, ok := h[key]; ok {
		return hv
	}

	return []string{}
}

func HeaderKeyToLower(h http.Header) http.Header {
	lh := http.Header{}
	for name := range h {
		lh[strings.ToLower(name)] = h[name]
	}
	return lh
}
