package main

import (
	"github.com/arnarpall/seccy/internal/encrypt"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/server"
	"github.com/arnarpall/seccy/internal/store/file"
)

func main() {
	logger := log.New()
	defer logger.Sync()

	enc, err := encrypt.NewClient("my-key")
	if err != nil {
		logger.Fatal(err)
	}

	store, err := file.NewFileStore(enc, "/tmp/arnar.vault")
	if err != nil {
		logger.Fatal(err)
	}

	s := server.New(":4040", logger, store)
	if err := s.Serve(); err != nil {
		logger.Fatalw("Unable to start server", "error", err)
	}
}
