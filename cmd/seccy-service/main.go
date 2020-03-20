package main

import (
	"flag"

	"github.com/arnarpall/seccy/internal/encrypt"
	"github.com/arnarpall/seccy/internal/env"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/server"
	"github.com/arnarpall/seccy/internal/server/grpc"
	"github.com/arnarpall/seccy/internal/store/file"
)

type serverOptions struct {
	encryptionKey string
	storePath     string
	listenAddress string
}

var opts = new(serverOptions)

func main() {
	flag.StringVar(&opts.encryptionKey, "encryption-key", env.GetEnvOrDefault("ENCRYPTION_KEY", ""), "The encryption key to use")
	flag.StringVar(&opts.storePath, "store-path", env.GetEnvOrDefault("STORE_PATH", ""), "The path to the data store")
	flag.StringVar(&opts.listenAddress, "listen-address", env.GetEnvOrDefault("LISTEN_ADDRESS", server.DefaultServerAddress), "The address to listen on")
	flag.Parse()

	logger := log.New()
	defer logger.Sync()

	if opts.encryptionKey == "" {
		logger.Fatal("Encryption key is missing")
	}
	if opts.storePath == "" {
		logger.Fatal("Store path is missing")
	}

	enc, err := encrypt.NewClient(opts.encryptionKey)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infow("Using an encrypted filestore", "path", opts.storePath)
	store, err := file.NewFileStore(enc, opts.storePath)
	if err != nil {
		logger.Fatal(err)
	}

	s := grpc.New(opts.listenAddress, logger, store)
	if err := s.Serve(); err != nil {
		logger.Fatalw("Unable to start server", "error", err)
	}
}
