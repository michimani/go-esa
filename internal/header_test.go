package internal_test

import (
	"net/http"
	"testing"

	"github.com/michimani/go-esa/v2/internal"
	"github.com/stretchr/testify/assert"
)

func Test_HeaderValue(t *testing.T) {
	cases := []struct {
		name   string
		header http.Header
		key    string
		expect []string
	}{
		{
			name: "normal",
			header: http.Header{
				"key1": []string{"value1-1", "value1-2"},
				"key2": []string{"value2-1", "value2-2"},
			},
			key:    "key1",
			expect: []string{"value1-1", "value1-2"},
		},
		{
			name: "normal: not exists key",
			header: http.Header{
				"key1": []string{"value1-1", "value1-2"},
				"key2": []string{"value2-1", "value2-2"},
			},
			key:    "key0",
			expect: []string{},
		},
		{
			name:   "normal: empty header",
			header: http.Header{},
			key:    "key",
			expect: []string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			vs := internal.HeaderValues(c.key, c.header)

			assert.Len(tt, vs, len(c.expect))
			assert.Equal(tt, c.expect, vs)
		})
	}
}

func Test_HeaderKeyToLower(t *testing.T) {
	cases := []struct {
		name   string
		header http.Header
		expect http.Header
	}{
		{
			name: "ok",
			header: http.Header{
				"Key1": []string{"value1-1", "value1-2"},
				"KEY2": []string{"value2-1", "value2-2"},
				"key3": []string{"value3-1", "value3-2"},
			},
			expect: http.Header{
				"key1": []string{"value1-1", "value1-2"},
				"key2": []string{"value2-1", "value2-2"},
				"key3": []string{"value3-1", "value3-2"},
			},
		},
		{
			name:   "ok: empty",
			header: http.Header{},
			expect: http.Header{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			lh := internal.HeaderKeyToLower(c.header)

			assert.Len(tt, lh, len(c.expect))
			assert.Equal(tt, c.expect, lh)
		})
	}
}
