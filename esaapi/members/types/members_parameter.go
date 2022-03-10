package types

import (
	"io"
	"net/url"
	"strings"

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

func (p *MembersGetParam) Body() (io.Reader, error) {
	return nil, nil
}

var membersGetParamQueryParams = map[string]struct{}{
	"sort":     {},
	"order":    {},
	"page":     {},
	"per_page": {},
}

func (p *MembersGetParam) ResolveEndpoint(endpointBase string) string {
	if p == nil {
		return ""
	}

	if p.TeamName == "" {
		return ""
	}

	encoded := url.QueryEscape(p.TeamName)
	endpoint := strings.Replace(endpointBase, ":team_name", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := internal.QueryString(pm, membersGetParamQueryParams)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *MembersGetParam) ParameterMap() map[string]string {
	m := internal.GeneratePaginationParamsMap(p, nil)

	if p.Sort.IsValid() {
		m["sort"] = string(p.Sort)
	}

	if p.Order.IsValid() {
		m["order"] = string(p.Order)
	}

	return m
}

type MembersScreenNameDeleteParam struct {
	TeamName   string
	ScreenName string
}

func (p *MembersScreenNameDeleteParam) Body() (io.Reader, error) {
	return nil, nil
}

func (p *MembersScreenNameDeleteParam) ResolveEndpoint(endpointBase string) string {
	if p == nil {
		return ""
	}

	if p.TeamName == "" || p.ScreenName == "" {
		return ""
	}

	encodedTeamName := url.QueryEscape(p.TeamName)
	encodedScreenName := url.QueryEscape(p.ScreenName)
	endpoint := strings.Replace(endpointBase, ":team_name", encodedTeamName, 1)
	endpoint = strings.Replace(endpoint, ":screen_name", encodedScreenName, 1)

	return endpoint
}

func (p *MembersScreenNameDeleteParam) ParameterMap() map[string]string {
	return nil
}
