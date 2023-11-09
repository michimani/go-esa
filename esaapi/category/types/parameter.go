package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/michimani/go-esa/v2/internal"
)

// BatchMoveInput is struct for the parameter for
// POST v1/teams/:team_name/categories/batch_move
type BatchMoveInput struct {
	TeamName string `json:"-"` // required

	From string `json:"from"` // required
	To   string `json:"to"`   // required
}

func (p *BatchMoveInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.From == "" || p.To == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "BatchMoveInput.TeamName, BatchMoveInput.From, BatchMoveInput.To")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  strings.NewReader(string(json)),
	}, nil
}
