package gesa_test

import (
	"testing"

	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_EsaAPIVersion_IsValid(t *testing.T) {
	cases := []struct {
		name   string
		v      gesa.EsaAPIVersion
		expect bool
	}{
		{
			name:   "valid: v1",
			v:      gesa.EsaAPIVersionV1,
			expect: true,
		},
		{
			name:   "invalid: v2",
			v:      gesa.EsaAPIVersion("v2"),
			expect: false,
		},
		{
			name:   "invalid: other string",
			v:      gesa.EsaAPIVersion("other"),
			expect: false,
		},
		{
			name:   "invalid: empty",
			v:      gesa.EsaAPIVersion(""),
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			b := c.v.IsValid()
			asst.Equal(c.expect, b)
		})
	}
}

func Test_EsaAPIVersion_String(t *testing.T) {
	cases := []struct {
		name   string
		v      gesa.EsaAPIVersion
		expect string
	}{
		{
			name:   "ok",
			v:      gesa.EsaAPIVersion("test-version"),
			expect: "test-version",
		},
		{
			name:   "ok: empty",
			v:      gesa.EsaAPIVersion(""),
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			s := c.v.String()
			asst.Equal(c.expect, s)
		})
	}
}

func Test_EsaAPIVersion_IsEmpty(t *testing.T) {
	cases := []struct {
		name   string
		v      gesa.EsaAPIVersion
		expect bool
	}{
		{
			name:   "not empty",
			v:      gesa.EsaAPIVersion("test-version"),
			expect: false,
		},
		{
			name:   "empty",
			v:      gesa.EsaAPIVersion(""),
			expect: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			s := c.v.IsEmpty()
			asst.Equal(c.expect, s)
		})
	}
}

func Test_EsaAPIVersion_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name    string
		v       gesa.EsaAPIVersion
		base    string
		expect  string
		wantErr bool
	}{
		{
			name:   "ok",
			v:      gesa.EsaAPIVersionV1,
			base:   "/:esa_api_version/hoge/endpoint",
			expect: "/v1/hoge/endpoint",
		},
		{
			name:   "ok: not versioned endpoint",
			v:      gesa.EsaAPIVersionV1,
			base:   "/hoge/endpoint",
			expect: "/hoge/endpoint",
		},
		{
			name:    "error: invalid api version",
			v:       gesa.EsaAPIVersion("v2"),
			base:    "/hoge/endpoint",
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			e, err := c.v.ResolveEndpoint(c.base)

			if c.wantErr {
				asst.NotNil(err)
				asst.Empty(e)
				return
			}
			asst.Nil(err)
			asst.Equal(c.expect, e)
		})
	}
}
