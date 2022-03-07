package internal_test

import (
	"testing"

	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_QueryValue(t *testing.T) {
	cases := []struct {
		name   string
		params []string
		expect string
	}{
		{
			name:   "normal",
			params: []string{"param1", "param2", "param3"},
			expect: "param1,param2,param3",
		},
		{
			name:   "normal: only one param",
			params: []string{"param1"},
			expect: "param1",
		},
		{
			name:   "normal: empty params",
			params: []string{},
			expect: "",
		},
		{
			name:   "normal: nil params",
			params: nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			q := internal.QueryValue(c.params)
			assert.Equal(tt, c.expect, q)
		})
	}
}

func Test_QueryString(t *testing.T) {
	cases := []struct {
		name     string
		includes map[string]struct{}
		params   map[string]string
		expect   string
	}{
		{
			name: "ok",
			includes: map[string]struct{}{
				"key1": {},
			},
			params: map[string]string{
				"key1": "value1",
			},
			expect: "key1=value1",
		},
		{
			name: "ok: some params",
			includes: map[string]struct{}{
				"key1": {},
				"key2": {},
			},
			params: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			expect: "key1=value1&key2=value2",
		},
		{
			name:     "ok: empty includes",
			includes: map[string]struct{}{},
			params: map[string]string{
				"key1": "value1",
			},
			expect: "",
		},
		{
			name: "ok: empty params",
			includes: map[string]struct{}{
				"key1": {},
			},
			params: map[string]string{},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			q := internal.QueryString(c.params, c.includes)
			assert.Equal(tt, c.expect, q)
		})
	}
}

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

func Test_GeneratePaginationParamsMap(t *testing.T) {
	cases := []struct {
		name   string
		p      *testParameter
		m      map[string]string
		expect map[string]string
	}{
		{
			name:   "both not null",
			p:      &testParameter{Page: gesa.NewPageNumber(1), PerPage: gesa.NewPageNumber(2)},
			m:      map[string]string{"hoge": "hogevalue"},
			expect: map[string]string{"hoge": "hogevalue", "page": "1", "per_page": "2"},
		},
		{
			name:   "both not null: map nil",
			p:      &testParameter{Page: gesa.NewPageNumber(1), PerPage: gesa.NewPageNumber(2)},
			m:      nil,
			expect: map[string]string{"page": "1", "per_page": "2"},
		},
		{
			name:   "page",
			p:      &testParameter{Page: gesa.NewPageNumber(1)},
			expect: map[string]string{"page": "1"},
		},
		{
			name:   "perPage",
			p:      &testParameter{PerPage: gesa.NewPageNumber(2)},
			expect: map[string]string{"per_page": "2"},
		},
		{
			name:   "both null",
			p:      &testParameter{},
			expect: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			am := internal.GeneratePaginationParamsMap(c.p, c.m)
			asst.Equal(c.expect, am)
		})
	}
}
