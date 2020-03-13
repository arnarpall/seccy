package server

import (
	"context"
	"fmt"
	"net"

	"github.com/arnarpall/seccy/api/proto/seccy"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/store"
	"github.com/arnarpall/seccy/internal/version"
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

func (s *seccyServer) Set(ctx context.Context, req *seccy.SetRequest) (*seccy.Empty, error) {
	s.logger.Infof("Setting key %s to value %s", req.Key, req.Value)
	err := s.store.Set(req.Key, req.Value)
	if err != nil {
		s.logger.Errorf("Unable to set value %s for key %s %s, %v", req.Value, req.Key, err)
	}

	return &seccy.Empty{}, err
}

func (s *seccyServer) Get(ctx context.Context, req *seccy.GetRequest) (*seccy.GetResponse, error) {
	s.logger.Infof("Getting value for key %s", req.Key)
	val, err := s.store.Get(req.Key)
	if err != nil {
		s.logger.Errorf("Unable to get value for key ting value for key %s, %v", req.Key, err)
		return nil, err
	}

	return &seccy.GetResponse{
		Value: val,
	}, nil
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
