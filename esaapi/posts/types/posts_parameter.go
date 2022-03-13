package types

import (
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
)

type PostsGetSort string

const (
	PostsGetSortUpdated   PostsGetSort = "updated" // default
	PostsGetSortCreated   PostsGetSort = "created"
	PostsGetSortNumber    PostsGetSort = "number"
	PostsGetSortStars     PostsGetSort = "stars"
	PostsGetSortWatches   PostsGetSort = "watches"
	PostsGetSortComments  PostsGetSort = "comments"
	PostsGetSortBestMatch PostsGetSort = "best_match"
)

func (s PostsGetSort) IsValid() bool {
	return s == PostsGetSortUpdated ||
		s == PostsGetSortCreated ||
		s == PostsGetSortNumber ||
		s == PostsGetSortStars ||
		s == PostsGetSortWatches ||
		s == PostsGetSortComments ||
		s == PostsGetSortBestMatch
}

type PostsGetOrder string

const (
	PostsGetOrderDesc PostsGetOrder = "desc" // default
	PostsGetOrderAsc  PostsGetOrder = "asc"
)

func (o PostsGetOrder) IsValid() bool {
	return o == PostsGetOrderAsc || o == PostsGetOrderDesc
}

type PostsGetParam struct {
	// Path parameter
	TeamName string

	// Query parameters
	Q       string
	Include string
	Sort    PostsGetSort
	Order   PostsGetOrder

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *PostsGetParam) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *PostsGetParam) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *PostsGetParam) EsaAPIParameter() *internal.EsaAPIParameter {
	if p == nil {
		return nil
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	qp := internal.QueryParameterList{}
	if p.Q != "" {
		qp = append(qp, internal.QueryParameter{Key: "q", Value: p.Q})
	}
	if p.Include != "" {
		qp = append(qp, internal.QueryParameter{Key: "include", Value: p.Include})
	}
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
