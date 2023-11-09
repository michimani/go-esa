package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/v2/esaapi/comment/types"
	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
	"github.com/stretchr/testify/assert"
)

func Test_ListPostCommentsInput_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListPostCommentsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListPostCommentsInput{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListPostCommentsInput{}, 0, false},
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

func Test_ListPostCommentsInput_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListPostCommentsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListPostCommentsInput{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListPostCommentsInput{}, 0, false},
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

func Test_ListPostCommentsInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListPostCommentsInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListPostCommentsInput{
				TeamName:   "test-team",
				PostNumber: 1,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "with page",
			p: &types.ListPostCommentsInput{
				TeamName:   "test-team",
				PostNumber: 1,
				Page:       gesa.NewPageNumber(1),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
				},
			},
		},
		{
			name: "with per_page",
			p: &types.ListPostCommentsInput{
				TeamName:   "test-team",
				PostNumber: 1,
				PerPage:    gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "with all",
			p: &types.ListPostCommentsInput{
				TeamName:   "test-team",
				PostNumber: 1,
				Page:       gesa.NewPageNumber(1),
				PerPage:    gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.ListPostCommentsInput{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only PostNumber",
			p: &types.ListPostCommentsInput{
				PostNumber: 1,
				Page:       gesa.NewPageNumber(1),
				PerPage:    gesa.NewPageNumber(2),
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

func Test_GetCommentInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.GetCommentInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.GetCommentInput{
				TeamName:  "test-team",
				CommentID: 1,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.GetCommentInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only CommentID",
			p: &types.GetCommentInput{
				CommentID: 1,
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

func Test_CreateCommentInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CreateCommentInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CreateCommentInput{
				TeamName:   "test-team",
				PostNumber: 1,
				BodyMD:     "test comment",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"comment":{"body_md":"test comment"}}`),
			},
		},
		{
			name: "ok: has user",
			p: &types.CreateCommentInput{
				TeamName:   "test-team",
				PostNumber: 1,
				BodyMD:     "test comment",
				User:       gesa.String("test-user"),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"comment":{"body_md":"test comment","user":"test-user"}}`),
			},
		},
		{
			name: "ng: not has required parameter: team_name is empty",
			p: &types.CreateCommentInput{
				PostNumber: 1,
				BodyMD:     "test comment",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.CreateCommentInput{
				TeamName: "test-team",
				BodyMD:   "test comment",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: body_md is empty",
			p: &types.CreateCommentInput{
				TeamName:   "test-team",
				PostNumber: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: all empty",
			p:       &types.CreateCommentInput{},
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
			ep, err := c.p.EsaAPIParameter()
			asst := assert.New(tt)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			asst.Equal(c.expect, ep)
		})
	}
}

func Test_UpdateCommentInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.UpdateCommentInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.UpdateCommentInput{
				TeamName:  "test-team",
				CommentID: 1,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"comment":{}}`),
			},
		},
		{
			name: "ok: has mody_md",
			p: &types.UpdateCommentInput{
				TeamName:  "test-team",
				CommentID: 1,
				BodyMD:    gesa.String("test comment (updated)"),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"comment":{"body_md":"test comment (updated)"}}`),
			},
		},
		{
			name: "ok: has user",
			p: &types.UpdateCommentInput{
				TeamName:  "test-team",
				CommentID: 1,
				BodyMD:    gesa.String("test comment (updated)"),
				User:      gesa.String("test-user"),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"comment":{"body_md":"test comment (updated)","user":"test-user"}}`),
			},
		},
		{
			name: "ng: not has required parameter: team_name is empty",
			p: &types.UpdateCommentInput{
				CommentID: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.UpdateCommentInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: all empty",
			p:       &types.UpdateCommentInput{},
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
			ep, err := c.p.EsaAPIParameter()
			asst := assert.New(tt)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			asst.Equal(c.expect, ep)
		})
	}
}

func Test_DeleteCommentInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.DeleteCommentInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.DeleteCommentInput{
				TeamName:  "test-team",
				CommentID: 1,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.DeleteCommentInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only CommentID",
			p: &types.DeleteCommentInput{
				CommentID: 1,
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

func Test_ListTeamCommentsInput_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListTeamCommentsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListTeamCommentsInput{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListTeamCommentsInput{}, 0, false},
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

func Test_ListTeamCommentsInput_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListTeamCommentsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListTeamCommentsInput{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListTeamCommentsInput{}, 0, false},
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

func Test_ListTeamCommentsInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListTeamCommentsInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListTeamCommentsInput{
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
			p: &types.ListTeamCommentsInput{
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
			p: &types.ListTeamCommentsInput{
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
			p: &types.ListTeamCommentsInput{
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
			p: &types.ListTeamCommentsInput{
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
