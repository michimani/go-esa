package gesa_test

import (
	"testing"

	"github.com/michimani/go-esa/v2/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	cases := []struct {
		name string
		s    string
	}{
		{
			name: "normal",
			s:    "test string",
		},
		{
			name: "normal: empty string",
			s:    "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			sp := gesa.String(c.s)
			assert.Equal(tt, c.s, *sp)
		})
	}
}

func Test_StringValue(t *testing.T) {
	cases := []struct {
		name string
		s    string
	}{
		{
			name: "normal",
			s:    "test string",
		},
		{
			name: "normal: empty string",
			s:    "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			sv := gesa.StringValue(&c.s)
			assert.Equal(tt, c.s, sv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		sv := gesa.StringValue(nil)
		assert.Empty(tt, sv)
	})
}

func Test_Bool(t *testing.T) {
	cases := []struct {
		name string
		b    bool
	}{
		{
			name: "normal: true",
			b:    true,
		},
		{
			name: "normal: false",
			b:    false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			bp := gesa.Bool(c.b)
			assert.Equal(tt, c.b, *bp)
		})
	}
}

func Test_BoolValue(t *testing.T) {
	cases := []struct {
		name string
		b    bool
	}{
		{
			name: "normal: true",
			b:    true,
		},
		{
			name: "normal: false",
			b:    false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			bv := gesa.BoolValue(&c.b)
			assert.Equal(tt, c.b, bv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		bv := gesa.BoolValue(nil)
		assert.False(tt, bv)
	})
}

func Test_Int(t *testing.T) {
	cases := []struct {
		name string
		i    int
	}{
		{
			name: "normal: zero",
			i:    0,
		},
		{
			name: "normal: positive",
			i:    1,
		},
		{
			name: "normal: negative",
			i:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ip := gesa.Int(c.i)
			assert.Equal(tt, c.i, *ip)
		})
	}
}

func Test_IntValue(t *testing.T) {
	cases := []struct {
		name string
		i    int
	}{
		{
			name: "normal: zero",
			i:    0,
		},
		{
			name: "normal: positive",
			i:    1,
		},
		{
			name: "normal: negative",
			i:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			iv := gesa.IntValue(&c.i)
			assert.Equal(tt, c.i, iv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		iv := gesa.IntValue(nil)
		assert.Equal(tt, iv, 0)
	})
}
