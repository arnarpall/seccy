package client

import (
	"context"
	"fmt"
	"io"

	"github.com/arnarpall/seccy/api/proto/seccy"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type Setter interface {
	Set(key, val string) error
}

type Getter interface {
	Get(key string) (string, error)
}

type Lister interface {
	ListKeys() (<-chan string, error)
}

type Client interface {
	Getter
	Setter
	Lister
}

type client struct {
	seccy  seccy.SeccyClient
	logger *log.Logger
}

func (c *client) Get(key string) (string, error) {
	c.logger.Debugw("Getting value", "key", key)
	r := &seccy.GetRequest{
		Key: key,
	}
	resp, err := c.seccy.Get(context.TODO(), r)
	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

func (c *client) Set(key, val string) error {
	c.logger.Debugw("Setting key value", "key", key, "value", val)
	r := &seccy.SetRequest{
		Key:   key,
		Value: val,
	}

	_, err := c.seccy.Set(context.TODO(), r)
	return err
}

func (c *client) ListKeys() (<-chan string, error) {
	c.logger.Debug("Listing all keys")
	keys, err := c.seccy.ListKeys(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}

	ch := make(chan string, 1)
	go func(ch chan<- string, receiver seccy.Seccy_ListKeysClient) {
		for {
			in, err := keys.Recv()
			if err != nil || err == io.EOF {
				close(ch)
				return
			}

			ch <- in.Key
		}
	}(ch, keys)

	return ch, nil
}

func New(address string, logger *log.Logger) (Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to server %s, %w", address, err)
	}

	c := seccy.NewSeccyClient(conn)
	return &client{
		seccy:  c,
		logger: logger,
	}, nil
}
