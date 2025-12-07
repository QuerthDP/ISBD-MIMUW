package tests

import (
	"context"
	"testing"

	"github.com/smogork/ISBD-MIMUW/pit"
)

func TableCreationTest(t *testing.T) {
	apiClient := pit.DbClient(pit.BaseURL)
	ctx := context.Background()

	t.Run("TableEmptyList", func(t *testing.T) {
		apiClient.SchemaAPI.ListTables(ctx).Execute()
	})

	t.Run("TableCreationAndList", func(t *testing.T) {
		// Implement table creation tests here
	})

	t.Run("TableCreationAndDetails", func(t *testing.T) {
		// Implement table creation tests here
	})

	t.Run("TableDoubleCreation", func(t *testing.T) {
		// Implement table creation tests here
	})

	t.Run("TableDoubleRemove", func(t *testing.T) {
		// Implement table creation tests here
	})

}
