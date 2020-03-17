package store

import "errors"

var ErrKeyNotFound = errors.New("key not found")

type Store interface {
	Set(key, val string) error
	Get(key string) (string, error)
	ListKeys() ([]string, error)
}
