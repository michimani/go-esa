package internal_test

import (
	"testing"

	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
	"github.com/stretchr/testify/assert"
)

type testParameter struct {
	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *testParameter) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *testParameter) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func Test_GeneratePaginationParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *testParameter
		expect internal.QueryParameterList
	}{
		{
			name: "both not null",
			p:    &testParameter{Page: gesa.NewPageNumber(1), PerPage: gesa.NewPageNumber(2)},
			expect: internal.QueryParameterList{
				{Key: "page", Value: "1"},
				{Key: "per_page", Value: "2"},
			},
		},
		{
			name: "page",
			p:    &testParameter{Page: gesa.NewPageNumber(1)},
			expect: internal.QueryParameterList{
				{Key: "page", Value: "1"},
			},
		},
		{
			name: "perPage",
			p:    &testParameter{PerPage: gesa.NewPageNumber(2)},
			expect: internal.QueryParameterList{
				{Key: "per_page", Value: "2"},
			},
		},
		{
			name:   "both null",
			p:      &testParameter{},
			expect: internal.QueryParameterList{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			am := internal.GeneratePaginationParameter(c.p)
			asst.Equal(c.expect, am)
		})
	}
}

func Test_QueryParameterList_QueryString(t *testing.T) {
	cases := []struct {
		name   string
		qps    internal.QueryParameterList
		expect string
	}{
		{
			name:   "ok: empty",
			qps:    internal.QueryParameterList{},
			expect: "",
		},
		{
			name: "ok: one parameter",
			qps: internal.QueryParameterList{
				{Key: "test-key", Value: "test-value"},
			},
			expect: "?test-key=test-value",
		},
		{
			name: "ok: two parameters",
			qps: internal.QueryParameterList{
				{Key: "test-key1", Value: "test-value1"},
				{Key: "test-key2", Value: "test-value2"},
			},
			expect: "?test-key1=test-value1&test-key2=test-value2",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			qs := c.qps.QueryString()
			asst.Equal(c.expect, qs)
		})
	}
}
