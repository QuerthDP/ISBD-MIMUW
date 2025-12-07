package tests

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/smogork/ISBD-MIMUW/pit"
)

var BaseURL string

// Command-line flags for test configuration
var (
	dbImage     = flag.String("db-image", "", "Docker image name (env: DB_IMAGE, default: isbd-mimuw-db:latest)")
	dbHostname  = flag.String("db-hostname", "", "Hostname of running database (env: DB_HOSTNAME, default: localhost)")
	dbPort      = flag.String("db-port", "", "Port on which database listens (env: DB_PORT, default: 8080)")
	dbRunDocker = flag.String("db-run-docker", "", "Skip docker container and use existing database (env: DB_RUN_DOCKER, default: false)")
)

func applyFlagToEnv() {
	// Set environment variables from flags if provided
	// Flags take precedence over environment variables
	if *dbImage != "" {
		os.Setenv("DB_IMAGE", *dbImage)
	}
	if *dbHostname != "" {
		os.Setenv("DB_HOSTNAME", *dbHostname)
	}
	if *dbPort != "" {
		os.Setenv("DB_PORT", *dbPort)
	}
	if *dbRunDocker != "" {
		os.Setenv("DB_RUN_DOCKER", *dbRunDocker)
	}
}

func TestMain(m *testing.M) {
	// Apply flags to environment variables
	applyFlagToEnv()

	ctx := context.Background()
	base, teardown, err := pit.StartTestContainer(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to start test container:", err)
		os.Exit(1)
	}
	BaseURL = base

	code := m.Run()

	if teardown != nil {
		teardown()
	}

	os.Exit(code)
}
