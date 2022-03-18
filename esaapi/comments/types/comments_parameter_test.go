package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/esaapi/comments/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_CommentsGetParam_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.CommentsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.CommentsGetParam{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.CommentsGetParam{}, 0, false},
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

func Test_CommentsGetParam_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.CommentsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.CommentsGetParam{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.CommentsGetParam{}, 0, false},
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

func Test_CommentsGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CommentsGetParam
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CommentsGetParam{
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
			p: &types.CommentsGetParam{
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
			p: &types.CommentsGetParam{
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
			p: &types.CommentsGetParam{
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
			p: &types.CommentsGetParam{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only PostNumber",
			p: &types.CommentsGetParam{
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

func Test_CommentsCommentIDGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CommentsCommentIDGetParam
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CommentsCommentIDGetParam{
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
			p: &types.CommentsCommentIDGetParam{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only CommentID",
			p: &types.CommentsCommentIDGetParam{
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

func Test_CommentsPostParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CommentsPostParam
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CommentsPostParam{
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
			p: &types.CommentsPostParam{
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
			p: &types.CommentsPostParam{
				PostNumber: 1,
				BodyMD:     "test comment",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.CommentsPostParam{
				TeamName: "test-team",
				BodyMD:   "test comment",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: body_md is empty",
			p: &types.CommentsPostParam{
				TeamName:   "test-team",
				PostNumber: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: all empty",
			p:       &types.CommentsPostParam{},
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

func Test_CommentsCommentIDPatchParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CommentsCommentIDPatchParam
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CommentsCommentIDPatchParam{
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
			p: &types.CommentsCommentIDPatchParam{
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
			p: &types.CommentsCommentIDPatchParam{
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
			p: &types.CommentsCommentIDPatchParam{
				CommentID: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.CommentsCommentIDPatchParam{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: all empty",
			p:       &types.CommentsCommentIDPatchParam{},
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

func Test_CommentsCommentIDDeleteParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CommentsCommentIDDeleteParam
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CommentsCommentIDDeleteParam{
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
			p: &types.CommentsCommentIDDeleteParam{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only CommentID",
			p: &types.CommentsCommentIDDeleteParam{
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

func Test_CommentsTeamNameGetParam_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.CommentsTeamNameGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.CommentsTeamNameGetParam{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.CommentsTeamNameGetParam{}, 0, false},
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

func Test_CommentsTeamNameGetParam_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.CommentsTeamNameGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.CommentsTeamNameGetParam{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.CommentsTeamNameGetParam{}, 0, false},
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

func Test_CommentsTeamNameGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CommentsTeamNameGetParam
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CommentsTeamNameGetParam{
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
			p: &types.CommentsTeamNameGetParam{
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
			p: &types.CommentsTeamNameGetParam{
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
			p: &types.CommentsTeamNameGetParam{
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
			p: &types.CommentsTeamNameGetParam{
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
