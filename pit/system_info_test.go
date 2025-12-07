package pit_tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type SystemInformation struct {
	InterfaceVersion string `json:"interfaceVersion"`
	Version          string `json:"version"`
	Author           string `json:"author"`
	Uptime           int64  `json:"uptime"`
}

var container tc.Container
var baseURL string

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
		baseURL = fmt.Sprintf("http://%s:%s", host, mp.Port())
	} else {
		baseURL = "http://localhost:8080"
	}

	code := m.Run()

	if container != nil {
		_ = container.Terminate(ctx)
	}

	os.Exit(code)
}

func TestSystemInfo(t *testing.T) {
	resp, err := http.Get(baseURL + "/system/info")
	if err != nil {
		t.Fatalf("GET /system/info failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("unexpected status %d: %s", resp.StatusCode, string(body))
	}
	var si SystemInformation
	if err := json.NewDecoder(resp.Body).Decode(&si); err != nil {
		t.Fatalf("decoding response failed: %v", err)
	}
	t.Logf("system version: %s", si.Version)
	if si.Version == "" {
		t.Fatalf("version field empty")
	}
}
