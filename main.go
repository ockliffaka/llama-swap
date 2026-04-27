package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mostlygeek/llama-swap/proxy"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var (
		configFile  = flag.String("config", "config.yaml", "path to configuration file")
		listenAddr  = flag.String("listen", "127.0.0.1:8080", "address to listen on")
		showVersion = flag.Bool("version", false, "print version information and exit")
		logRequests = flag.Bool("log-requests", false, "log all incoming requests")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("llama-swap version %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	// Validate config file exists before proceeding
	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", *configFile)
	}

	log.Printf("llama-swap %s starting up", version)
	log.Printf("loading configuration from: %s", *configFile)

	cfg, err := proxy.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server, err := proxy.NewServer(cfg, proxy.ServerOptions{
		ListenAddr:  *listenAddr,
		LogRequests: *logRequests,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Printf("listening on %s", *listenAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
