package pit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	apiclient "github.com/smogork/ISBD-MIMUW/pit/client"
)

var BaseURL string

func DbClient(url string) *apiclient.APIClient {
	// Create and configure API client
	cfg := apiclient.NewConfiguration()
	cfg.Servers[0].URL = url
	cfg.HTTPClient = &http.Client{}
	return apiclient.NewAPIClient(cfg)
}

func waitForHTTP(url string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	client := &http.Client{}
	for time.Now().Before(deadline) {
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
		resp, err := client.Do(req)
		if err == nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return nil
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("timeout waiting for %s", url)
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	var container tc.Container

	if os.Getenv("SKIP_DOCKER") == "" {
		image := os.Getenv("DB_IMAGE")
		if image == "" {
			image = "isbd-mimuw-db:latest"
		}

		req := tc.ContainerRequest{
			Image:        image,
			ExposedPorts: []string{"8080/tcp"},
			WaitingFor:   wait.ForHTTP("/system/info").WithPort("8080/tcp").WithStartupTimeout(30 * time.Second),
		}

		cont, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{ContainerRequest: req, Started: true})
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to start container:", err)
			os.Exit(1)
		}
		container = cont

		host, err := cont.Host(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to get container host:", err)
			_ = cont.Terminate(ctx)
			os.Exit(1)
		}
		mp, err := cont.MappedPort(ctx, "8080/tcp")
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to get mapped port:", err)
			_ = cont.Terminate(ctx)
			os.Exit(1)
		}
		BaseURL = fmt.Sprintf("http://%s:%s", host, mp.Port())
	} else {
		BaseURL = "http://localhost:8080"
	}

	// Wait for system info to be available
	// Do not use client, because of potentially invalid responses
	waitForHTTP(BaseURL+"/system/info", 30*time.Second)

	code := m.Run()

	if container != nil {
		_ = container.Terminate(ctx)
	}

	os.Exit(code)
}
