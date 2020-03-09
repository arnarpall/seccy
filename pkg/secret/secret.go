package secret

type Vault interface {
	Set(key, val string) error
	Get(key string) (string, error)
}