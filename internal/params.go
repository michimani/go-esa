package internal

import (
	"net/url"
	"strconv"
	"strings"
)

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
