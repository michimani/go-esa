package internal

import (
	"io"
	"net/url"
	"strconv"
)

type QueryParameter struct {
	Key   string
	Value string
}

type QueryParameterList []QueryParameter

func (qpl QueryParameterList) QueryString() string {
	q := url.Values{}
	for _, qp := range qpl {
		q.Add(qp.Key, qp.Value)
	}

	qs := q.Encode()
	if qs == "" {
		return ""
	}

	return "?" + qs
}

type PathParameter struct {
	Key   string
	Value string
}

type PathParameterList []PathParameter

type EsaAPIParameter struct {
	Query QueryParameterList
	Path  PathParameterList
	Body  io.Reader
}

func GeneratePaginationParameter(p IPaginationParameters) QueryParameterList {
	qp := QueryParameterList{}

	if page, ok := p.PageValue(); ok && page > 0 {
		qp = append(qp, QueryParameter{Key: "page", Value: strconv.Itoa(page)})
	}

	if perPage, ok := p.PerPageValue(); ok && perPage > 0 {
		qp = append(qp, QueryParameter{Key: "per_page", Value: strconv.Itoa(perPage)})
	}

	return qp
}
