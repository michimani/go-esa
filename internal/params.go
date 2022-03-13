package internal

import (
	"io"
	"net/url"
	"strconv"
	"strings"
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

func QueryValue(params []string) string {
	if len(params) == 0 {
		return ""
	}

	return strings.Join(params, ",")
}

func QueryString(paramsMap map[string]string, includes map[string]struct{}) string {
	q := url.Values{}
	for k, v := range paramsMap {
		if _, ok := includes[k]; ok {
			q.Add(k, v)
		}
	}

	return q.Encode()
}

func GeneratePaginationParamsMap(p IPaginationParameters, paramsMap map[string]string) map[string]string {
	if paramsMap == nil {
		paramsMap = map[string]string{}
	}

	if page, ok := p.PageValue(); ok && page > 0 {
		paramsMap["page"] = strconv.Itoa(page)
	}

	if perPage, ok := p.PerPageValue(); ok && perPage > 0 {
		paramsMap["per_page"] = strconv.Itoa(perPage)
	}

	return paramsMap
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
