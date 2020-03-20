package main

import (
	"flag"

	"github.com/arnarpall/seccy/internal/env"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/server"
	"github.com/arnarpall/seccy/internal/server/rest"
	"github.com/arnarpall/seccy/pkg/client"
)

const DefaultListenAddress = ":8080"

type serverOptions struct {
	listenAddress string
	seccyAddress  string
}

var opts = new(serverOptions)

func main() {
	flag.StringVar(&opts.listenAddress, "listen-address", env.GetEnvOrDefault("LISTEN_ADDRESS", DefaultListenAddress), "The address to listen on")
	flag.StringVar(&opts.seccyAddress, "seccy-server", env.GetEnvOrDefault("SECCY_SERVER", server.DefaultServerAddress), "The address of the seccy server")
	flag.Parse()

	logger := log.New()

	c, err := client.New(opts.seccyAddress, logger)
	if err != nil {
		logger.Fatalf("Unable to connect to seccy server, make sure that the server has been started", "error", err)
	}

	s := rest.New(logger, opts.listenAddress, c)
	if err := s.Serve(); err != nil {
		logger.Fatalw("unable to start server", "error", err)
	}
}
