package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/esaapi/category/types"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_BatchMoveInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.BatchMoveInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.BatchMoveInput{
				TeamName: "test-team",
				From:     "test-from",
				To:       "test-to",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"from":"test-from","to":"test-to"}`),
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.BatchMoveInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only From",
			p: &types.BatchMoveInput{
				From: "test-from",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only To",
			p: &types.BatchMoveInput{
				To: "test-to",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter",
			p:       &types.BatchMoveInput{},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: nil",
			p:       nil,
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			ep, err := c.p.EsaAPIParameter()
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			assert.Equal(tt, c.expect, ep)
		})
	}
}
