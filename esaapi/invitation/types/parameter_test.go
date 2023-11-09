package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/v2/esaapi/invitation/types"
	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
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

func Test_ListEmailInvitationsInput_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListEmailInvitationsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListEmailInvitationsInput{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListEmailInvitationsInput{}, 0, false},
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

func Test_ListEmailInvitationsInput_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListEmailInvitationsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListEmailInvitationsInput{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListEmailInvitationsInput{}, 0, false},
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

func Test_ListEmailInvitationsInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListEmailInvitationsInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListEmailInvitationsInput{
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
			name: "with page",
			p: &types.ListEmailInvitationsInput{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
				},
			},
		},
		{
			name: "with per_page",
			p: &types.ListEmailInvitationsInput{
				TeamName: "test-team",
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "with all",
			p: &types.ListEmailInvitationsInput{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "ng: not has required parameter",
			p: &types.ListEmailInvitationsInput{
				Page:    gesa.NewPageNumber(1),
				PerPage: gesa.NewPageNumber(2),
			},
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

func Test_DeleteEmailInvitationInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.DeleteEmailInvitationInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.DeleteEmailInvitationInput{
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
			p: &types.DeleteEmailInvitationInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only Code",
			p: &types.DeleteEmailInvitationInput{
				Code: "test-code",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter",
			p:       &types.DeleteEmailInvitationInput{},
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
