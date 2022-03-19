package types

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
)

// PostNumberStargazersGetParam is struct for the parameter for
// GET /v1/teams/:team_name/posts/:post_number/stargazers
type PostNumberStargazersGetParam struct {
	TeamName   string // required
	PostNumber int    // required

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *PostNumberStargazersGetParam) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *PostNumberStargazersGetParam) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *PostNumberStargazersGetParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "PostNumberStargazersGetParam.TeamName, PostNumberStargazersGetParam.PostNumber")
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
