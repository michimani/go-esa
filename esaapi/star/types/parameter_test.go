package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/v2/esaapi/star/types"
	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
	"github.com/stretchr/testify/assert"
)

func Test_ListPostStargazersInput_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListPostStargazersInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListPostStargazersInput{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListPostStargazersInput{}, 0, false},
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

func Test_ListPostStargazersInput_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListPostStargazersInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListPostStargazersInput{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListPostStargazersInput{}, 0, false},
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

func Test_ListPostStargazersInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListPostStargazersInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListPostStargazersInput{
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
			p: &types.ListPostStargazersInput{
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
			p: &types.ListPostStargazersInput{
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
			p: &types.ListPostStargazersInput{
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
			p: &types.ListPostStargazersInput{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only PostNumber",
			p: &types.ListPostStargazersInput{
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

func Test_CreatePostStarInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CreatePostStarInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CreatePostStarInput{
				TeamName:   "test-team",
				PostNumber: 1,
				Body:       "test body",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"body":"test body"}`),
			},
		},
		{
			name: "ng: not has required parameter: team_name is empty",
			p: &types.CreatePostStarInput{
				PostNumber: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.CreatePostStarInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: body is empty",
			p: &types.CreatePostStarInput{
				TeamName:   "test-team",
				PostNumber: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: both are empty",
			p:       &types.CreatePostStarInput{},
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

func Test_DeletePostStarInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.DeletePostStarInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.DeletePostStarInput{
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
			name: "ng: not has required parameter: has only TeamName",
			p: &types.DeletePostStarInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only PostNumber",
			p: &types.DeletePostStarInput{
				PostNumber: 1,
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

func Test_ListCommentStargazersInput_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListCommentStargazersInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListCommentStargazersInput{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListCommentStargazersInput{}, 0, false},
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

func Test_ListCommentStargazersInput_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListCommentStargazersInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListCommentStargazersInput{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListCommentStargazersInput{}, 0, false},
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

func Test_ListCommentStargazersInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListCommentStargazersInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListCommentStargazersInput{
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
			name: "with page",
			p: &types.ListCommentStargazersInput{
				TeamName:  "test-team",
				CommentID: 1,
				Page:      gesa.NewPageNumber(1),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
				},
			},
		},
		{
			name: "with per_page",
			p: &types.ListCommentStargazersInput{
				TeamName:  "test-team",
				CommentID: 1,
				PerPage:   gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "with all",
			p: &types.ListCommentStargazersInput{
				TeamName:  "test-team",
				CommentID: 1,
				Page:      gesa.NewPageNumber(1),
				PerPage:   gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "ng: not has required parameter: has only TeamName",
			p: &types.ListCommentStargazersInput{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only CommentID",
			p: &types.ListCommentStargazersInput{
				CommentID: 1,
				Page:      gesa.NewPageNumber(1),
				PerPage:   gesa.NewPageNumber(2),
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

func Test_CreateCommentStarInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CreateCommentStarInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CreateCommentStarInput{
				TeamName:  "test-team",
				CommentID: 1,
				Body:      "test body",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":comment_id", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"body":"test body"}`),
			},
		},
		{
			name: "ng: not has required parameter: team_name is empty",
			p: &types.CreateCommentStarInput{
				CommentID: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: comment_id is empty",
			p: &types.CreateCommentStarInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: body is empty",
			p: &types.CreateCommentStarInput{
				TeamName:  "test-team",
				CommentID: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: both are empty",
			p:       &types.CreateCommentStarInput{},
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

func Test_DeleteCommentStarInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.DeleteCommentStarInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.DeleteCommentStarInput{
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
			p: &types.DeleteCommentStarInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: has only CommentID",
			p: &types.DeleteCommentStarInput{
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
