package internal

import (
	"net/url"
	"strings"
)

func ResolveEndpoint(base string, pp PathParameterList, qp QueryParameterList) string {
	endpoint := base
	for _, p := range pp {
		endpoint = strings.Replace(endpoint, p.Key, url.QueryEscape(p.Value), 1)
	}

	return endpoint + qp.QueryString()
}
