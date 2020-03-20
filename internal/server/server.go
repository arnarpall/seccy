package server

import "errors"

const DefaultServerAddress = ":4040"

var (
	EmptyKey = errors.New("empty key")
	EmptyValue = errors.New("empty value")
)

type Server interface {
	Serve() error
}
