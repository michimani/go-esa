package types

import (
	"errors"
	"fmt"

	"github.com/michimani/go-esa/v2/internal"
)

type GetStatsInput struct {
	TeamName string
}

func (p *GetStatsInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "GetStatsInput.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}
