package types

import "io"

type OAuthTokenInfoGetParam struct{}

func (p *OAuthTokenInfoGetParam) Body() (io.Reader, error) {
	return nil, nil
}

func (p *OAuthTokenInfoGetParam) ResolveEndpoint(endpointBase string) string {
	return endpointBase
}

func (p *OAuthTokenInfoGetParam) ParameterMap() map[string]string {
	return nil
}
