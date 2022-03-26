package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/esaapi/invitation/types"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_GetURLInvitationInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.GetURLInvitationInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.GetURLInvitationInput{
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
			name:    "ng: not has required parameter",
			p:       &types.GetURLInvitationInput{},
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

func Test_RegenerateURLInvitationInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.RegenerateURLInvitationInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.RegenerateURLInvitationInput{
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
			name:    "ng: not has required parameter",
			p:       &types.RegenerateURLInvitationInput{},
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

func Test_CreateEmailInvitationsInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CreateEmailInvitationsInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CreateEmailInvitationsInput{
				TeamName: "test-team",
				Emails:   []string{"e1@example.com", "e2@example.com"},
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"member":{"emails":["e1@example.com","e2@example.com"]}}`),
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.CreateEmailInvitationsInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only emails",
			p: &types.CreateEmailInvitationsInput{
				Emails: []string{"e1@example.com", "e2@example.com"},
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter",
			p:       &types.CreateEmailInvitationsInput{},
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
