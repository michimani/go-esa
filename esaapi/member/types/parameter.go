package types

import (
	"errors"
	"fmt"

	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
)

type ListMembersSort string

const (
	ListMembersSortPostsCount   ListMembersSort = "posts_count" // default
	ListMembersSortJoined       ListMembersSort = "joined"
	ListMembersSortLastAccessed ListMembersSort = "last_accessed"
)

func (s ListMembersSort) IsValid() bool {
	return s == ListMembersSortPostsCount || s == ListMembersSortJoined || s == ListMembersSortLastAccessed
}

type ListMembersOrder string

const (
	ListMembersOrderDesc ListMembersOrder = "desc" // default
	ListMembersOrderAsc  ListMembersOrder = "asc"
)

func (o ListMembersOrder) IsValid() bool {
	return o == ListMembersOrderAsc || o == ListMembersOrderDesc
}

type ListMembersInput struct {
	TeamName string

	Sort  ListMembersSort
	Order ListMembersOrder

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *ListMembersInput) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *ListMembersInput) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *ListMembersInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "ListMembersInput.TeamName")
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
	}, nil
}

type DeleteMemberInput struct {
	TeamName          string
	ScreenNameOrEmail string
}

func (p *DeleteMemberInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.ScreenNameOrEmail == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "DeleteMemberInput.TeamName, DeleteMemberInput.ScreenNameOrEmail")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":screen_name_or_email", Value: p.ScreenNameOrEmail})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}
