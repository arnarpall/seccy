package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/arnarpall/seccy/internal/encrypt"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/store/file"
	"github.com/arnarpall/seccy/pkg/seccy"
	"github.com/arnarpall/seccy/pkg/vault"
)

type seccyServer struct {
	logger *log.Logger
	vault  *vault.Client
}

func (s *seccyServer) Set(ctx context.Context, req *seccy.SetRequest) (*seccy.Empty, error) {
	s.logger.Infof("Setting key %s to value %s", req.Key, req.Value)
	err := s.vault.Set(req.Key, req.Value)
	if err != nil {
		s.logger.Errorf("Unable to set value %s for key %s %s, %v", req.Value, req.Key, err)
	}

	return &seccy.Empty{}, err
}

func (s *seccyServer) Get(ctx context.Context, req *seccy.GetRequest) (*seccy.GetResponse, error) {
	s.logger.Infof("Getting value for key %s", req.Key)
	val, err := s.vault.Get(req.Key)
	if err != nil {
		s.logger.Errorf("Unable to get value for key ting value for key %s, %v", req.Key, err)
		return nil, err
	}

	return &seccy.GetResponse{
		Value: val,
	}, nil
}

func main() {
	logger := log.New()
	defer logger.Sync()

	enc, err := encrypt.NoOp("my-key")
	if err != nil {
		logger.Fatal(err)
	}

	store, err := file.NewFileStore(enc, "/tmp/arnar.vault")
	if err != nil {
		logger.Fatal(err)
	}

	client := vault.NewClient(store)
	server := seccyServer{
		logger: logger,
		vault: client,
	}

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic("Unable to listen on port 4040")
	}

	gs := grpc.NewServer()
	seccy.RegisterSeccyServer(gs, &server)
	if err := gs.Serve(lis); err != nil {
		panic("Unable to start grcp server")
	}
}
