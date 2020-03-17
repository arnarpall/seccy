package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/arnarpall/seccy/api/proto/seccy"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/store"
	"github.com/arnarpall/seccy/internal/version"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type seccyServer struct {
	address string
	logger  *log.Logger
	store   store.Store
}

type Server interface {
	Serve() error
}

func New(address string, logger *log.Logger, store store.Store) Server {
	return &seccyServer{
		address: address,
		logger:  logger,
		store:   store,
	}
}

func (s *seccyServer) Set(_ context.Context, req *seccy.SetRequest) (*empty.Empty, error) {
	s.logger.Infow("Setting key value", "key", req.Key, "value", req.Value)
	err := s.store.Set(req.Key, req.Value)
	if err != nil {
		s.logger.Errorw("Unable to set key value", "key", req.Key, "value", req.Value, "error", err)
		return &empty.Empty{}, errors.New("unable to connect to store")
	}

	return &empty.Empty{}, err
}

func (s *seccyServer) Get(_ context.Context, req *seccy.GetRequest) (*seccy.GetResponse, error) {
	s.logger.Infow("Getting value", "key", req.Key)
	val, err := s.store.Get(req.Key)
	if err != nil {
		s.logger.Errorw("Unable to get value for key", "key", req.Key, "error", err)

		if errors.Is(err, store.ErrKeyNotFound) {
			return nil, err
		}
	}

	return &seccy.GetResponse{
		Value: val,
	}, nil
}

func (s *seccyServer) ListKeys(_ *empty.Empty, stream seccy.Seccy_ListKeysServer) error {
	s.logger.Info("Listing all keys")
	keys, err := s.store.ListKeys()
	if err != nil {
		s.logger.Errorw("Unable to list all keys", "error", err)
		return err
	}

	s.logger.Infof("found %d keys", len(keys))

	for _, k := range keys {
		s.logger.Debugw("sending key", "key", k)
		if err := stream.Send(&seccy.KeyResponse{Key: k}); err != nil {
			return err
		}
	}

	return nil
}

func (s *seccyServer) Serve() error {
	s.logger.Infow("Starting server",
		"version", version.BuildVersion,
		"buildDate", version.BuildDate,
		"address", s.address)

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("unable to listen on address %s, %w", s.address, err)
	}

	gs := grpc.NewServer()
	seccy.RegisterSeccyServer(gs, s)
	if err := gs.Serve(lis); err != nil {
		return fmt.Errorf("error serving grpc connection, %w", err)
	}

	return nil
}
