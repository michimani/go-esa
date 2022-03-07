package types_test

import (
	"io"
	"testing"

	"github.com/michimani/go-esa/esaapi/teams/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_TeamsGetParam_Body(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.TeamsGetParam
		want    io.Reader
		wantErr bool
	}{
		{"ok, nil", nil, nil, false},
		{"ok: empty", &types.TeamsGetParam{}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			r, err := c.p.Body()
			if c.wantErr {
				asst.NotNil(err)
				asst.Nil(r)
				return
			}

			asst.Nil(err)
			asst.Equal(c.want, r)
		})
	}
}

func Test_TeamsGetParam_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.TeamsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.TeamsGetParam{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.TeamsGetParam{}, 0, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			i, b := c.p.PageValue()
			asst.Equal(c.expectInt, i)
			asst.Equal(c.expectBool, b)
		})
	}
}

func Test_TeamsGetParam_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.TeamsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.TeamsGetParam{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.TeamsGetParam{}, 0, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			i, b := c.p.PerPageValue()
			asst.Equal(c.expectInt, i)
			asst.Equal(c.expectBool, b)
		})
	}
}
