package internal

type Store interface {
	Set(key, val string) error
	Get(key string) (string, error)
}