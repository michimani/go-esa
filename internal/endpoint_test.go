package internal_test

import (
	"testing"

	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name   string
		base   string
		pp     internal.PathParameterList
		qp     internal.QueryParameterList
		expect string
	}{
		{
			name: "ok: not need path parameter, no query parameter",
			base: "test-endpoint",
			pp: internal.PathParameterList{
				{Key: ":path-key", Value: "path-value"},
			},
			qp:     internal.QueryParameterList{},
			expect: "test-endpoint",
		},
		{
			name: "ok: not need path parameter, some query parameters",
			base: "test-endpoint",
			pp: internal.PathParameterList{
				{Key: ":path-key", Value: "path-value"},
			},
			qp: internal.QueryParameterList{
				{Key: "query-key1", Value: "query-value1"},
				{Key: "query-key2", Value: "query-value2"},
			},
			expect: "test-endpoint?query-key1=query-value1&query-key2=query-value2",
		},
		{
			name: "ok: need path parameter, no query parameter",
			base: "test-endpoint/:path-key/test",
			pp: internal.PathParameterList{
				{Key: ":path-key", Value: "path-value"},
			},
			qp:     internal.QueryParameterList{},
			expect: "test-endpoint/path-value/test",
		},
		{
			name: "ok: some need path parameter, no query parameter",
			base: "test-endpoint/:path-key1/:path-key2/test",
			pp: internal.PathParameterList{
				{Key: ":path-key1", Value: "path-value1"},
				{Key: ":path-key2", Value: "path-value2"},
			},
			qp: internal.QueryParameterList{
				{Key: "query-key1", Value: "query-value1"},
				{Key: "query-key2", Value: "query-value2"},
			},
			expect: "test-endpoint/path-value1/path-value2/test?query-key1=query-value1&query-key2=query-value2",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			ep := internal.ResolveEndpoint(c.base, c.pp, c.qp)
			asst.Equal(c.expect, ep)
		})
	}
}
