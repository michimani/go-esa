package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/michimani/go-esa/internal"
)

type ListEmojisInput struct {
	TeamName string // required
}

func (p *ListEmojisInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "ListEmojisInput.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}

type CreateEmojiInput struct {
	TeamName string // required

	Code       string  // required
	OriginCode *string // required if you add the other alias to already exists emoji
	Image      *string // required if you create new emoji, BASE64 encoded string
}

type createEmojiPayload struct {
	Emoji createEmojiPayloadEmoji `json:"emoji"`
}

type createEmojiPayloadEmoji struct {
	Code       string  `json:"code"`
	OriginCode *string `json:"origin_code,omitempty"`
	Image      *string `json:"image,omitempty"`
}

func (p *CreateEmojiInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.Code == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreateEmojiInput.TeamName, CreateEmojiInput.Code")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	payload := &createEmojiPayload{
		Emoji: createEmojiPayloadEmoji{
			Code:       p.Code,
			OriginCode: p.OriginCode,
			Image:      p.Image,
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
