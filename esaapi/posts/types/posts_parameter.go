package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
)

type PostsGetSort string

const (
	PostsGetSortUpdated   PostsGetSort = "updated" // default
	PostsGetSortCreated   PostsGetSort = "created"
	PostsGetSortNumber    PostsGetSort = "number"
	PostsGetSortStars     PostsGetSort = "stars"
	PostsGetSortWatches   PostsGetSort = "watches"
	PostsGetSortComments  PostsGetSort = "comments"
	PostsGetSortBestMatch PostsGetSort = "best_match"
)

func (s PostsGetSort) IsValid() bool {
	return s == PostsGetSortUpdated ||
		s == PostsGetSortCreated ||
		s == PostsGetSortNumber ||
		s == PostsGetSortStars ||
		s == PostsGetSortWatches ||
		s == PostsGetSortComments ||
		s == PostsGetSortBestMatch
}

type PostsGetOrder string

const (
	PostsGetOrderDesc PostsGetOrder = "desc" // default
	PostsGetOrderAsc  PostsGetOrder = "asc"
)

func (o PostsGetOrder) IsValid() bool {
	return o == PostsGetOrderAsc || o == PostsGetOrderDesc
}

type PostsGetParam struct {
	// Path parameter
	TeamName string

	// Query parameters
	Q       string
	Include string
	Sort    PostsGetSort
	Order   PostsGetOrder

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *PostsGetParam) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *PostsGetParam) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *PostsGetParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "PostsGetParam.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	qp := internal.QueryParameterList{}
	if p.Q != "" {
		qp = append(qp, internal.QueryParameter{Key: "q", Value: p.Q})
	}
	if p.Include != "" {
		qp = append(qp, internal.QueryParameter{Key: "include", Value: p.Include})
	}
	if p.Sort.IsValid() {
		qp = append(qp, internal.QueryParameter{Key: "sort", Value: string(p.Sort)})
	}
	if p.Order.IsValid() {
		qp = append(qp, internal.QueryParameter{Key: "order", Value: string(p.Order)})
	}

	pagination := internal.GeneratePaginationParameter(p)
	qp = append(qp, pagination...)

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: qp,
		Body:  nil,
	}, nil
}

type PostsPostNumberGetParam struct {
	// Path parameter
	TeamName   string
	PostNumber int

	// Query parameters
	Include string
}

func (p *PostsPostNumberGetParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "PostsPostNumberGetParam.TeamName, PostsPostNumberGetParam.PostNumber")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":post_number", Value: strconv.Itoa(p.PostNumber)})

	qp := internal.QueryParameterList{}
	if p.Include != "" {
		qp = append(qp, internal.QueryParameter{Key: "include", Value: p.Include})
	}

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: qp,
		Body:  nil,
	}, nil
}

type PostsPostParam struct {
	// Path parameter
	TeamName string `json:"-"`

	// Payload
	Name     string // required
	BodyMD   *string
	Tags     []*string
	Category *string
	Wip      *bool
	Message  *string
	User     *string
}

type postsPostPayload struct {
	Post postsPostPayloadPost `json:"post"`
}

type postsPostPayloadPost struct {
	Name     string    `json:"name"` // required
	BodyMD   *string   `json:"body_md,omitempty"`
	Tags     []*string `json:"tags,omitempty"`
	Category *string   `json:"category,omitempty"`
	Wip      *bool     `json:"wip,omitempty"`
	Message  *string   `json:"message,omitempty"`
	User     *string   `json:"user,omitempty"`
}

func (p *PostsPostParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "PostsPostParam.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	if p.Name == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "PostsPostParam.Name")
	}

	payload := &postsPostPayload{
		Post: postsPostPayloadPost{
			Name:     p.Name,
			BodyMD:   p.BodyMD,
			Tags:     p.Tags,
			Category: p.Category,
			Wip:      p.Wip,
			Message:  p.Message,
			User:     p.User,
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
