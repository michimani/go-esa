package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
)

// ListPostCommentsInput is struct for the parameter for
// GET /v1/teams/:team_name/posts/:post_number/comments
type ListPostCommentsInput struct {
	TeamName   string
	PostNumber int

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *ListPostCommentsInput) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *ListPostCommentsInput) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *ListPostCommentsInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "ListPostCommentsInput.TeamName, ListPostCommentsInput.PostNumber")
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

// GetCommentInput is struct for the parameter for
// GET /v1/teams/:team_name/comments/:comment_id
type GetCommentInput struct {
	TeamName  string
	CommentID int
}

func (p *GetCommentInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.CommentID == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "GetCommentInput.TeamName, GetCommentInput.CommentID")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":comment_id", Value: strconv.Itoa(p.CommentID)})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}

// CreateCommentInput is struct for the parameter for
// POST /v1/teams/:team_name/posts/:post_number/comments
type CreateCommentInput struct {
	// Path parameter
	TeamName   string
	PostNumber int

	// Payload
	BodyMD string // required
	User   *string
}

type createCommentPayload struct {
	Comment createCommentPayloadComment `json:"comment"`
}

type createCommentPayloadComment struct {
	BodyMD string  `json:"body_md"`
	User   *string `json:"user,omitempty"`
}

func (p *CreateCommentInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreateCommentInput.TeamName, CreateCommentInput.PostNumber")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":post_number", Value: strconv.Itoa(p.PostNumber)})

	if p.BodyMD == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreateCommentInput.BodyMD")
	}

	payload := &createCommentPayload{
		Comment: createCommentPayloadComment{
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

// UpdateCommentInput is struct for the parameter for
// PATCH /v1/teams/:team_name/comments/:comment_id
type UpdateCommentInput struct {
	// Path parameter
	TeamName  string
	CommentID int

	// Payload
	BodyMD *string
	User   *string
}

type updateCommentPayload struct {
	Comment updateCommentPayloadComment `json:"comment"`
}

type updateCommentPayloadComment struct {
	BodyMD *string `json:"body_md,omitempty"`
	User   *string `json:"user,omitempty"`
}

func (p *UpdateCommentInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.CommentID == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "UpdateCommentInput.TeamName, UpdateCommentInput.CommentID")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":comment_id", Value: strconv.Itoa(p.CommentID)})

	payload := &updateCommentPayload{
		Comment: updateCommentPayloadComment{
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

// DeleteCommentInput is struct for the parameter for
// DELETE /v1/teams/:team_name/comments/:comment_id
type DeleteCommentInput struct {
	TeamName  string
	CommentID int
}

func (p *DeleteCommentInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.CommentID == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "DeleteCommentInput.TeamName, DeleteCommentInput.CommentID")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":comment_id", Value: strconv.Itoa(p.CommentID)})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}

// ListTeamCommentsInput is struct for the parameter for
// GET /v1/teams/:team_name/comments
type ListTeamCommentsInput struct {
	TeamName string

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *ListTeamCommentsInput) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *ListTeamCommentsInput) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *ListTeamCommentsInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "ListTeamCommentsInput.TeamName")
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
