package configs

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// ParseFlags parses command-line flags and returns the configuration file path.
// It defaults to "./config.yml" if no --config flag is provided.
// Returns an error if the config path is invalid.
func ParseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}

// ValidateConfigPath validates that the given path exists and is a file (not a directory).
// Returns an error if the path does not exist or is a directory.
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("invalid file format")
	}
	return nil
}

// Run starts the HTTP server with graceful shutdown support.
// It listens for SIGINT and SIGTERM signals and performs a graceful shutdown
// with a 5-second timeout when a termination signal is received.
// The server configuration (host, port, timeouts) is read from the Config struct.
func (c *Config) Run(r *mux.Router) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:         c.Configs.Server.Host + ":" + c.Configs.Server.Port,
		Handler:      r,
		ReadTimeout:  c.Configs.Server.Timeout.Read * time.Second,
		WriteTimeout: c.Configs.Server.Timeout.Write * time.Second,
		IdleTimeout:  c.Configs.Server.Timeout.Idle * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
