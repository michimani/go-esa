package types

import (
	"errors"

	"github.com/michimani/go-esa/v2/internal"
)

type GetMeInput struct {
	Include string
}

func (p *GetMeInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	qp := internal.QueryParameterList{}
	if p.Include != "" {
		qp = append(qp, internal.QueryParameter{Key: "include", Value: p.Include})
	}

	return &internal.EsaAPIParameter{
		Path:  internal.PathParameterList{},
		Query: qp,
		Body:  nil,
	}, nil
}
