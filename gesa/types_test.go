package gesa_test

import (
	"testing"
	"time"

	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_Timestamp_Time(t *testing.T) {
	ts := gesa.Timestamp(1461218696)
	tm := time.Unix(int64(1461218696), 0)

	cases := []struct {
		name string
		ts   *gesa.Timestamp
		want *time.Time
	}{
		{
			name: "ok: not nil",
			ts:   &ts,
			want: &tm,
		},
		{
			name: "ok: nil",
			ts:   nil,
			want: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			atm := c.ts.Time()

			asst.Equal(atm, c.want)
		})
	}
}

func Test_Timestamp_SafeTimestamp(t *testing.T) {
	ts := gesa.Timestamp(1461218696)

	cases := []struct {
		name string
		ts   *gesa.Timestamp
		want int64
	}{
		{
			name: "ok: not nil",
			ts:   &ts,
			want: 1461218696,
		},
		{
			name: "ok: nil",
			ts:   nil,
			want: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			st := c.ts.SafeTimestamp()

			asst.Equal(st, c.want)
		})
	}
}
