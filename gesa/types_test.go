package gesa_test

import (
	"testing"
	"time"

	"github.com/michimani/go-esa/v2/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_Timestamp_Time(t *testing.T) {
	ts := gesa.Timestamp(1461218696)
	tm := time.Unix(int64(1461218696), 0)

	cases := []struct {
		name   string
		ts     *gesa.Timestamp
		expect *time.Time
	}{
		{"ok: not nil", &ts, &tm},
		{"ok: nil", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			atm := c.ts.Time()

			asst.Equal(atm, c.expect)
		})
	}
}

func Test_Timestamp_SafeTimestamp(t *testing.T) {
	ts := gesa.Timestamp(1461218696)

	cases := []struct {
		name   string
		ts     *gesa.Timestamp
		expect int64
	}{
		{"ok: not nil", &ts, 1461218696},
		{"ok: nil", nil, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			st := c.ts.SafeTimestamp()

			asst.Equal(st, c.expect)
		})
	}
}

func Test_NewPageNumber(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		expect gesa.PageNumber
	}{
		{"0", 0, gesa.PageNumber(0)},
		{"positive", 1, gesa.PageNumber(1)},
		{"negative", -1, gesa.PageNumber(0)},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			p := gesa.NewPageNumber(c.n)
			asst.Equal(c.expect, *p)
		})
	}
}

func Test_PageNumber_IsNull(t *testing.T) {
	cases := []struct {
		name   string
		p      *gesa.PageNumber
		expect bool
	}{
		{"null: nil", nil, true},
		{"null: 0", gesa.NewPageNumber(0), true},
		{"not null", gesa.NewPageNumber(1), false},
		{"not null: negative number", gesa.NewPageNumber(-1), true},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			b := c.p.IsNull()
			asst.Equal(c.expect, b)
		})
	}
}

func Test_PageNumber_SafeInt(t *testing.T) {
	cases := []struct {
		name   string
		p      *gesa.PageNumber
		expect int
	}{
		{"null: nil", nil, 0},
		{"null: 0", gesa.NewPageNumber(0), 0},
		{"not null", gesa.NewPageNumber(1), 1},
		{"not null: negative number", gesa.NewPageNumber(-1), 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			b := c.p.SafeInt()
			asst.Equal(c.expect, b)
		})
	}
}
