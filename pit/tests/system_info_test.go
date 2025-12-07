package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/smogork/ISBD-MIMUW/pit"
)

func SystemInfoTest(t *testing.T) {
	apiClient := pit.DbClient(pit.BaseURL)

	t.Run("SystemInfo", func(t *testing.T) {
		ctx := context.Background()
		sysInfo, resp, err := apiClient.MetadataAPI.GetSystemInfo(ctx).Execute()
		if err != nil {
			t.Fatalf("GetSystemInfo failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("unexpected status %d", resp.StatusCode)
		}
		if sysInfo == nil {
			t.Fatalf("system info is nil")
		}
		t.Logf("system version: %s", sysInfo.Version)
		if sysInfo.Version == "" {
			t.Fatalf("version field empty")
		}
	})
}
