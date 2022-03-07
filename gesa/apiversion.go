package gesa

import (
	"errors"
	"strings"
)

type EsaAPIVersion string

const (
	versionSubject string = ":esa_api_version"

	EsaAPIVersionV1 EsaAPIVersion = "v1"

	DefaultAPIVersion EsaAPIVersion = EsaAPIVersionV1
)

var availables = map[EsaAPIVersion]struct{}{
	EsaAPIVersionV1: {},
}

func (v EsaAPIVersion) String() string {
	return string(v)
}

func (v EsaAPIVersion) IsValid() bool {
	_, ok := availables[v]
	return ok
}

func (v EsaAPIVersion) IsEmpty() bool {
	return v.String() == ""
}

func (v EsaAPIVersion) ResolveEndpoint(base string) (string, error) {
	if !v.IsValid() {
		return "", errors.New("invalid esa api version")
	}

	return strings.Replace(base, versionSubject, v.String(), 1), nil
}
