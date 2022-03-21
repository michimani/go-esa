package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/esaapi/emoji/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_ListEmojisInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListEmojisInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListEmojisInput{
				TeamName: "test-team",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name:    "ng: has no required parameter",
			p:       &types.ListEmojisInput{},
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

func Test_CreateEmojiInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CreateEmojiInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CreateEmojiInput{
				TeamName: "test-team",
				Code:     "test-code",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"emoji":{"code":"test-code"}}`),
			},
		},
		{
			name: "ok: with origin_code",
			p: &types.CreateEmojiInput{
				TeamName:   "test-team",
				Code:       "test-code",
				OriginCode: gesa.String("test-origin-code"),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"emoji":{"code":"test-code","origin_code":"test-origin-code"}}`),
			},
		},
		{
			name: "ok: with image",
			p: &types.CreateEmojiInput{
				TeamName: "test-team",
				Code:     "test-code",
				Image:    gesa.String("test-image"),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"emoji":{"code":"test-code","image":"test-image"}}`),
			},
		},
		{
			name: "ok: with all",
			p: &types.CreateEmojiInput{
				TeamName:   "test-team",
				Code:       "test-code",
				OriginCode: gesa.String("test-origin-code"),
				Image:      gesa.String("test-image"),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"emoji":{"code":"test-code","origin_code":"test-origin-code","image":"test-image"}}`),
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.CreateEmojiInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only Code",
			p: &types.CreateEmojiInput{
				Code: "test-code",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter",
			p:       &types.CreateEmojiInput{},
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

func Test_DeleteEmojiInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.DeleteEmojiInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.DeleteEmojiInput{
				TeamName: "test-team",
				Code:     "test-code",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":code", Value: "test-code"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.DeleteEmojiInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only Code",
			p: &types.DeleteEmojiInput{
				Code: "test-code",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter",
			p:       &types.DeleteEmojiInput{},
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
