package secret

import (
	"github.com/arnarpall/seccy/internal"
)

type Setter interface {
	Set(key, val string) error
}

type Getter interface {
	Get(key string) (string, error)
}

type Client struct {
	store internal.Store
}

func NewClient(store internal.Store) *Client {
	return &Client{
		store: store,
	}
}

func (c *Client) Set(key, val string) error {
	return c.store.Set(key, val)
}

func (c *Client) Get(key string) (string, error) {
	return c.store.Get(key)
}