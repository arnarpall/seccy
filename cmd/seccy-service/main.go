package main

import (
	"flag"
	"os"

	"github.com/arnarpall/seccy/internal/encrypt"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/server"
	"github.com/arnarpall/seccy/internal/store/file"
)

const DefaultListenAddress = ":4040"

type serverOptions struct {
	encryptionKey string
	storePath     string
	listenAddress string
}

var opts = new(serverOptions)

func main() {
	flag.StringVar(&opts.encryptionKey, "encryption-key", getEnvOrDefault("ENCRYPTION_KEY", ""), "The encryption key to use")
	flag.StringVar(&opts.storePath, "store-path", getEnvOrDefault("STORE_PATH", ""), "The path to the data store")
	flag.StringVar(&opts.listenAddress, "listen-address", getEnvOrDefault("LISTEN_ADDRESS", DefaultListenAddress), "The address to listen on")
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

	store, err := file.NewFileStore(enc, opts.storePath)
	if err != nil {
		logger.Fatal(err)
	}

	s := server.New(opts.listenAddress, logger, store)
	if err := s.Serve(); err != nil {
		logger.Fatalw("Unable to start server", "error", err)
	}
}

func getEnvOrDefault(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}
