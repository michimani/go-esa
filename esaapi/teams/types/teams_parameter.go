package types

import (
	"io"

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
