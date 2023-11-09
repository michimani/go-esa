package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
)

type ListPostsSort string

const (
	ListPostsSortUpdated   ListPostsSort = "updated" // default
	ListPostsSortCreated   ListPostsSort = "created"
	ListPostsSortNumber    ListPostsSort = "number"
	ListPostsSortStars     ListPostsSort = "stars"
	ListPostsSortWatches   ListPostsSort = "watches"
	ListPostsSortComments  ListPostsSort = "comments"
	ListPostsSortBestMatch ListPostsSort = "best_match"
)

func (s ListPostsSort) IsValid() bool {
	return s == ListPostsSortUpdated ||
		s == ListPostsSortCreated ||
		s == ListPostsSortNumber ||
		s == ListPostsSortStars ||
		s == ListPostsSortWatches ||
		s == ListPostsSortComments ||
		s == ListPostsSortBestMatch
}

type ListPostsOrder string

const (
	ListPostsOrderDesc ListPostsOrder = "desc" // default
	ListPostsOrderAsc  ListPostsOrder = "asc"
)

func (o ListPostsOrder) IsValid() bool {
	return o == ListPostsOrderAsc || o == ListPostsOrderDesc
}

type ListPostsInput struct {
	// Path parameter
	TeamName string

	// Query parameters
	Q       string
	Include string
	Sort    ListPostsSort
	Order   ListPostsOrder

	Page    *gesa.PageNumber
	PerPage *gesa.PageNumber
}

func (p *ListPostsInput) PageValue() (int, bool) {
	if p.Page.IsNull() {
		return 0, false
	}
	return p.Page.SafeInt(), true
}

func (p *ListPostsInput) PerPageValue() (int, bool) {
	if p.PerPage.IsNull() {
		return 0, false
	}
	return p.PerPage.SafeInt(), true
}

func (p *ListPostsInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "ListPostsInput.TeamName")
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

type GetPostInput struct {
	// Path parameter
	TeamName   string
	PostNumber int

	// Query parameters
	Include string
}

func (p *GetPostInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "GetPostInput.TeamName, GetPostInput.PostNumber")
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

type CreatePostInput struct {
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

type createPostPayload struct {
	Post createPostPayloadPost `json:"post"`
}

type createPostPayloadPost struct {
	Name     string    `json:"name"` // required
	BodyMD   *string   `json:"body_md,omitempty"`
	Tags     []*string `json:"tags,omitempty"`
	Category *string   `json:"category,omitempty"`
	Wip      *bool     `json:"wip,omitempty"`
	Message  *string   `json:"message,omitempty"`
	User     *string   `json:"user,omitempty"`
}

func (p *CreatePostInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreatePostInput.TeamName")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})

	if p.Name == "" {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "CreatePostInput.Name")
	}

	payload := &createPostPayload{
		Post: createPostPayloadPost{
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

type UpdatePostInput struct {
	// Path parameter
	TeamName   string `json:"-"`
	PostNumber int    `json:"-"`

	// Payload
	Name             string // required
	BodyMD           *string
	Tags             []*string
	Category         *string
	Wip              *bool
	Message          *string
	CreatedBy        *string // screen_name, enabled only owner
	UpdatedBy        *string // screen_name, enabled only owner
	OriginalRevision *OriginalRevision
}

type OriginalRevision struct {
	BodyMD *string `json:"body_md,omitempty"`
	Number *int    `json:"number,omitempty"`
	User   *string `json:"user,omitempty"`
}

type updatePostPayload struct {
	Post updatePostPayloadPost `json:"post"`
}

type updatePostPayloadPost struct {
	Name             string            `json:"name,omitempty"`
	BodyMD           *string           `json:"body_md,omitempty"`
	Tags             []*string         `json:"tags,omitempty"`
	Category         *string           `json:"category,omitempty"`
	Wip              *bool             `json:"wip,omitempty"`
	Message          *string           `json:"message,omitempty"`
	CreatedBy        *string           `json:"created_by,omitempty"`
	UpdatedBy        *string           `json:"updated_by,omitempty"`
	OriginalRevision *OriginalRevision `json:"original_revision,omitempty"`
}

func (p *UpdatePostInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "UpdatePostInput.TeamName, UpdatePostInput.PostNumber")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":post_number", Value: strconv.Itoa(p.PostNumber)})

	payload := &updatePostPayload{
		Post: updatePostPayloadPost{
			Name:             p.Name,
			BodyMD:           p.BodyMD,
			Tags:             p.Tags,
			Category:         p.Category,
			Wip:              p.Wip,
			Message:          p.Message,
			CreatedBy:        p.CreatedBy,
			UpdatedBy:        p.UpdatedBy,
			OriginalRevision: p.OriginalRevision,
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

type DeletePostInput struct {
	// Path parameter
	TeamName   string `json:"-"`
	PostNumber int    `json:"-"`
}

func (p *DeletePostInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	if p == nil {
		return nil, errors.New(internal.ErrorParameterIsNil)
	}

	pp := internal.PathParameterList{}
	if p.TeamName == "" || p.PostNumber == 0 {
		return nil, fmt.Errorf(internal.ErrorRequiredParameterEmpty, "DeletePostInput.TeamName, DeletePostInput.PostNumber")
	}
	pp = append(pp, internal.PathParameter{Key: ":team_name", Value: p.TeamName})
	pp = append(pp, internal.PathParameter{Key: ":post_number", Value: strconv.Itoa(p.PostNumber)})

	return &internal.EsaAPIParameter{
		Path:  pp,
		Query: internal.QueryParameterList{},
		Body:  nil,
	}, nil
}
