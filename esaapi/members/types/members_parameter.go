package types

import (
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
)

type MembersGetSort string

const (
	MembersGetSortPostsCount   MembersGetSort = "posts_count" // default
	MembersGetSortJoined       MembersGetSort = "joined"
	MembersGetSortLastAccessed MembersGetSort = "last_accessed"
)

func (s MembersGetSort) IsValid() bool {
	return s == MembersGetSortPostsCount || s == MembersGetSortJoined || s == MembersGetSortLastAccessed
}

type MembersGetOrder string

const (
	MembersGetOrderDesc MembersGetOrder = "desc" // default
	MembersGetOrderAsc  MembersGetOrder = "asc"
)

func (o MembersGetOrder) IsValid() bool {
	return o == MembersGetOrderAsc || o == MembersGetOrderDesc
}

type MembersGetParam struct {
	TeamName string

	Sort  MembersGetSort
	Order MembersGetOrder

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *MembersGetParam) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *MembersGetParam) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *MembersGetParam) EsaAPIParameter() *internal.EsaAPIParameter {
	if p == nil {
		return nil
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	qp := internal.QueryParameterList{}
	if p.Sort.IsValid() {
		qp = append(qp, internal.QueryParameter{Key: "sort", Value: string(p.Sort)})
	}
	if p.Order.IsValid() {
		qp = append(qp, internal.QueryParameter{Key: "order", Value: string(p.Order)})
	}

	pagination := internal.GeneratePaginationParameter(p)
	qp = append(qp, pagination...)

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: qp,
		Body:  nil,
	}
}

type MembersScreenNameDeleteParam struct {
	TeamName   string
	ScreenName string
}

func (p *MembersScreenNameDeleteParam) EsaAPIParameter() *internal.EsaAPIParameter {
	if p == nil {
		return nil
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.ScreenName == "" {
		return nil
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":screen_name", Value: p.ScreenName})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}
}
