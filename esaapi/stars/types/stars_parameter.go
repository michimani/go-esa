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

// ListPostStargazersInput is struct for the parameter for
// GET /v1/teams/:team_name/posts/:post_number/stargazers
type ListPostStargazersInput struct {
	TeamName   string // required
	PostNumber int    // required

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *ListPostStargazersInput) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *ListPostStargazersInput) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *ListPostStargazersInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "ListPostStargazersInput.TeamName, ListPostStargazersInput.PostNumber")
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

type CreatePostStarInput struct {
	// Path parameter
	TeamName   string `json:"-"`
	PostNumber int    `json:"-"`

	// Payload
	Body string `json:"body"` // required
}

func (p *CreatePostStarInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreatePostStarInput.TeamName, CreatePostStarInput.PostNumber")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":post_number", Value: strconv.Itoa(p.PostNumber)})

	if p.Body == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreatePostStarInput.Body")
	}

	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  strings.NewReader(string(json)),
	}, nil
}
