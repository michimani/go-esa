package types

import (
	"io"
	"net/url"
	"strings"
)

type StatsGetParam struct {
	TeamName string
}

func (p *StatsGetParam) Body() (io.Reader, error) {
	return nil, nil
}

func (p *StatsGetParam) ResolveEndpoint(endpointBase string) string {
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

func (p *StatsGetParam) ParameterMap() map[string]string {
	return nil
}
