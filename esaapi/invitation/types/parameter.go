package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/michimani/go-esa/internal"
)

type GetURLInvitationInput struct {
	TeamName string // required
}

func (p *GetURLInvitationInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "GetURLInvitationInput.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}

type RegenerateURLInvitationInput struct {
	TeamName string // required
}

func (p *RegenerateURLInvitationInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "RegenerateURLInvitationInput.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}

type CreateEmailInvitationsInput struct {
	TeamName string // required

	Emails []string // required
}

type createEmailInvitationsPayload struct {
	Member createEmailInvitationsPayloadEmailInvitations `json:"member"`
}

type createEmailInvitationsPayloadEmailInvitations struct {
	Emails []string `json:"emails"`
}

func (p *CreateEmailInvitationsInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreateEmailInvitationsInput.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	if len(p.Emails) == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreateEmailInvitationsInput.Emails")
	}
	payload := &createEmailInvitationsPayload{
		Member: createEmailInvitationsPayloadEmailInvitations{
			Emails: p.Emails,
		},
	}

	json, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  strings.NewReader(string(json)),
	}, nil
}
