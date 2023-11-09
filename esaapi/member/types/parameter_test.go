package types_test

import (
	"testing"

	"github.com/michimani/go-esa/v2/esaapi/member/types"
	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
	"github.com/stretchr/testify/assert"
)

func Test_ListMembersSort_IsValid(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		expect bool
	}{
		{"ok", "joined", true},
		{"ok", "posts_count", true},
		{"ok", "last_accessed", true},
		{"ng", "unknown", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			asst.Equal(c.expect, types.ListMembersSort(c.s).IsValid())
		})
	}
}

func Test_ListMembersOrder_IsValid(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		expect bool
	}{
		{"ok", "asc", true},
		{"ok", "desc", true},
		{"ng", "unknown", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			asst.Equal(c.expect, types.ListMembersOrder(c.s).IsValid())
		})
	}
}

func Test_ListMembersInput_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListMembersInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListMembersInput{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListMembersInput{}, 0, false},
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

func Test_ListMembersInput_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListMembersInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListMembersInput{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListMembersInput{}, 0, false},
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

func Test_ListMembersInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListMembersInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListMembersInput{
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
			p: &types.ListMembersInput{
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
			p: &types.ListMembersInput{
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
			name: "with sort",
			p: &types.ListMembersInput{
				TeamName: "test-team",
				Sort:     types.ListMembersSortJoined,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "sort", Value: "joined"},
				},
			},
		},
		{
			name: "with order",
			p: &types.ListMembersInput{
				TeamName: "test-team",
				Order:    types.ListMembersOrderAsc,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "order", Value: "asc"},
				},
			},
		},
		{
			name: "with all",
			p: &types.ListMembersInput{
				TeamName: "test-team",
				Sort:     types.ListMembersSortPostsCount,
				Order:    types.ListMembersOrderDesc,
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "sort", Value: "posts_count"},
					{Key: "order", Value: "desc"},
					{Key: "page", Value: "1"},
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "ng: not has required parameter",
			p: &types.ListMembersInput{
				Sort:    types.ListMembersSortPostsCount,
				Order:   types.ListMembersOrderDesc,
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

func Test_DeleteMemberInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.DeleteMemberInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.DeleteMemberInput{
				TeamName:          "test-team",
				ScreenNameOrEmail: "test-screen-name",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":screen_name_or_email", Value: "test-screen-name"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "ng: not has required parameter: only team_name",
			p: &types.DeleteMemberInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: only screen_name_or_email",
			p: &types.DeleteMemberInput{
				ScreenNameOrEmail: "test-screen-name",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: both",
			p:       &types.DeleteMemberInput{},
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
