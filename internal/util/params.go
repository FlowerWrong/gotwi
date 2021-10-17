package util

import (
	"io"
	"strings"
)

type Parameters interface {
	SetAccessToken(token string)
	AccessToken() string
	ResolveEndpoint(endpointBase string) string
	Body() io.Reader
	ParameterMap() map[string]string
}

func QueryValue(params []string) string {
	if len(params) == 0 {
		return ""
	}

	return strings.Join(params, ",")
}
