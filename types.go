package gesa

import "io"

type IParameters interface {
	Body() (io.Reader, error)
}

type IResponse interface{}
