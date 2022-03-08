package types

import (
	"io"
	"net/url"
	"strings"

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

func (p *TeamsGetParam) Body() (io.Reader, error) {
	return nil, nil
}

var teamsGetParamQueryParams = map[string]struct{}{
	"page":     {},
	"per_page": {},
}

func (p *TeamsGetParam) ResolveEndpoint(endpointBase string) string {
	if p == nil {
		return ""
	}

	endpoint := endpointBase

	pm := p.ParameterMap()
	qs := internal.QueryString(pm, teamsGetParamQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TeamsGetParam) ParameterMap() map[string]string {
	return internal.GeneratePaginationParamsMap(p, nil)
}

type TeamsTeamNameGetParam struct {
	TeamName string
}

func (p *TeamsTeamNameGetParam) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TeamsTeamNameGetParam) ResolveEndpoint(endpointBase string) string {
	if p == nil {
		return ""
	}

	if p.TeamName == "" {
		return ""
	}

	encoded := url.QueryEscape(p.TeamName)
	endpoint := strings.Replace(endpointBase, ":team_name", encoded, 1)

	return endpoint
}

func (p *TeamsTeamNameGetParam) ParameterMap() map[string]string {
	return nil
}
