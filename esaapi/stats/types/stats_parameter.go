package types

import "github.com/michimani/go-esa/internal"

type StatsGetParam struct {
	TeamName string
}

func (p *StatsGetParam) EsaAPIParameter() *internal.EsaAPIParameter {
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
