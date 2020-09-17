package main

import (
	"flag"
	"fmt"
	"github.com/rafaelthomazi/qa/qa/service"
	"github.com/rafaelthomazi/qa/server/rest"
	"go.uber.org/zap"
	"os"
)

func main() {
	opts := parseServerOptions()
	if opts == nil {
		fmt.Println("Exiting due to flag errors")
		os.Exit(2)
	}

	logger, _ := zap.NewProduction()
	opts.Logger = logger
	opts.DAO.Logger = logger

	svc := service.NewService(opts)
	if svc == nil {
		fmt.Println("Failed to create service")
		os.Exit(2)
	}

	// Create a channel for errors
	errc := make(chan error)

	server := rest.NewServer(svc, opts.HTTPPort, logger)

	go func() {
		logger.Info("HTTP service started listening", zap.String("addr", server.Addr))
		errc <- server.ListenAndServe()
	}()

	// Log errors
	logger.Info("exit", zap.Error(<-errc))
}

func parseServerOptions() *service.Config {
	cfg := new(service.Config)

	flag.StringVar(&(cfg.HTTPPort), "httpPort", os.Getenv("HTTP_PORT"), "Server HTTP port")
	flag.StringVar(&(cfg.DAO.URI), "mongodbURI", os.Getenv("MONGODB_URI"), "MongoDB URI")
	flag.StringVar(&(cfg.DAO.Database), "mongodbDatabase", os.Getenv("MONGODB_DATABASE"), "MongoDB Database")

	flag.Parse()

	if len(cfg.HTTPPort) < 2 {
		fmt.Printf("INVALID -httpPort: %v", cfg.HTTPPort)
		return nil
	}

	cfg.HTTPPort = fmt.Sprintf(":%v", cfg.HTTPPort)

	return cfg
}
