package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/smogork/ISBD-MIMUW/pit"
	"github.com/stretchr/testify/assert"
)

func SystemInfoTest(t *testing.T) {
	apiClient := pit.DbClient(pit.BaseURL)
	assert := assert.New(t)

	t.Run("SystemInfo", func(t *testing.T) {
		ctx := context.Background()
		sysInfo, resp, err := apiClient.MetadataAPI.GetSystemInfo(ctx).Execute()

		assert.NoError(err, "GetSystemInfo should not return an error")
		assert.Equal(http.StatusOK, resp.StatusCode, "GetSystemInfo should return status 200")
		assert.NotNil(sysInfo, "SystemInformation should not be nil")
		assert.NotEmpty(sysInfo.Version, "Version field should not be empty")
		assert.NotEmpty(sysInfo.Author, "Author field should not be empty")
	})
}
