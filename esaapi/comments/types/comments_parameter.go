package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
)

// CommentsGetParam is struct for the parameter for
// GET /v1/teams/:team_name/posts/:post_number/comments
type CommentsGetParam struct {
	TeamName   string
	PostNumber int

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *CommentsGetParam) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *CommentsGetParam) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *CommentsGetParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CommentsGetParam.TeamName, CommentsGetParam.PostNumber")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":post_number", Value: strconv.Itoa(p.PostNumber)})

	qp := internal.QueryParameterList{}
	pagination := internal.GeneratePaginationParameter(p)
	qp = append(qp, pagination...)

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: qp,
		Body:  nil,
	}, nil
}

// CommentsCommentIDGetParam is struct for the parameter for
// GET /v1/teams/:team_name/comments/:comment_id
type CommentsCommentIDGetParam struct {
	TeamName  string
	CommentID int
}

func (p *CommentsCommentIDGetParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.CommentID == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CommentsCommentIDGetParam.TeamName, CommentsCommentIDGetParam.CommentID")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":comment_id", Value: strconv.Itoa(p.CommentID)})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}

// CommentsPostParam is struct for the parameter for
// POST /v1/teams/:team_name/posts/:post_number/comments
type CommentsPostParam struct {
	// Path parameter
	TeamName   string
	PostNumber int

	// Payload
	BodyMD string // required
	User   *string
}

type CommentsPostPayload struct {
	Comment CommentsPostPayloadComment `json:"comment"`
}

type CommentsPostPayloadComment struct {
	BodyMD string  `json:"body_md"`
	User   *string `json:"user,omitempty"`
}

func (p *CommentsPostParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CommentsPostParam.TeamName, CommentsPostParam.PostNumber")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":post_number", Value: strconv.Itoa(p.PostNumber)})

	if p.BodyMD == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CommentsPostParam.BodyMD")
	}

	payload := &CommentsPostPayload{
		Comment: CommentsPostPayloadComment{
			BodyMD: p.BodyMD,
			User:   p.User,
		},
	}

	json, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  strings.NewReader(string(json)),
	}, nil
}

// CommentsCommentIDPatchParam is struct for the parameter for
// PATCH /v1/teams/:team_name/comments/:comment_id
type CommentsCommentIDPatchParam struct {
	// Path parameter
	TeamName  string
	CommentID int

	// Payload
	BodyMD *string
	User   *string
}

type CommentsCommentIDPatchPayload struct {
	Comment CommentsCommentIDPatchPayloadComment `json:"comment"`
}

type CommentsCommentIDPatchPayloadComment struct {
	BodyMD *string `json:"body_md,omitempty"`
	User   *string `json:"user,omitempty"`
}

func (p *CommentsCommentIDPatchParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.CommentID == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CommentsCommentIDPatchParam.TeamName, CommentsCommentIDPatchParam.CommentID")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":comment_id", Value: strconv.Itoa(p.CommentID)})

	payload := &CommentsCommentIDPatchPayload{
		Comment: CommentsCommentIDPatchPayloadComment{
			BodyMD: p.BodyMD,
			User:   p.User,
		},
	}

	json, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  strings.NewReader(string(json)),
	}, nil
}

// CommentsCommentIDDeleteParam is struct for the parameter for
// DELETE /v1/teams/:team_name/comments/:comment_id
type CommentsCommentIDDeleteParam struct {
	TeamName  string
	CommentID int
}

func (p *CommentsCommentIDDeleteParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.CommentID == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CommentsCommentIDDeleteParam.TeamName, CommentsCommentIDDeleteParam.CommentID")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":comment_id", Value: strconv.Itoa(p.CommentID)})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}

// CommentsTeamNameGetParam is struct for the parameter for
// GET /v1/teams/:team_name/comments
type CommentsTeamNameGetParam struct {
	TeamName string

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *CommentsTeamNameGetParam) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *CommentsTeamNameGetParam) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *CommentsTeamNameGetParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CommentsTeamNameGetParam.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	qp := internal.QueryParameterList{}
	pagination := internal.GeneratePaginationParameter(p)
	qp = append(qp, pagination...)

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: qp,
		Body:  nil,
	}, nil
}
