package types

import (
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
)

type TeamsGetParam struct {
	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *TeamsGetParam) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *TeamsGetParam) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *TeamsGetParam) EsaAPIParameter() *internal.EsaAPIParameter {
	if p == nil {
		return nil
	}

	qp := internal.QueryParameterList{}
	pagination := internal.GeneratePaginationParameter(p)
	qp = append(qp, pagination...)

	return &internal.EsaAPIParameter{
		Path:  internal.PathParameterList{},
		Query: qp,
		Body:  nil,
	}
}

type TeamsTeamNameGetParam struct {
	TeamName string
}

func (p *TeamsTeamNameGetParam) EsaAPIParameter() *internal.EsaAPIParameter {
	if p == nil {
		return nil
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}
}
