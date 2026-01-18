package tests

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/smogork/ISBD-MIMUW/pit"
	openapi1 "github.com/smogork/ISBD-MIMUW/pit/client/openapi1"
	"github.com/stretchr/testify/require"
)

// ============================================================================
// COPY Operation Tests (Interface v1)
// ============================================================================

func TestV1_Copy(t *testing.T) {
	RequireInterfaceVersion(t, 1)

	dbClient := DbClientV1(BaseURL)
	ctx := context.Background()

	t.Run("Copy_Success", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "people")

		// Execute COPY
		LoadTestDataV1(t, dbClient, ctx, "people")

		// Verify data was loaded
		result := ExecuteSelectStarV1(t, dbClient, ctx, "people")
		rows := ParseQueryResultsV1(result)
		require.Len(t, rows, 5, "Should have 5 rows after COPY")
	})

	t.Run("Copy_ToNonExistentTable_Fails", func(t *testing.T) {
		queryId, resp, err := SubmitCopyQueryV1(dbClient, ctx,
			"/data/tables/people/data.csv",
			"non_existent_table_xyz",
			true)

		// May fail at submission (400) or during execution (FAILED)
		if err == nil && resp.StatusCode == http.StatusOK {
			query, waitErr := WaitForQueryCompletionV1(dbClient, ctx, queryId, 10*time.Second)
			require.NoError(t, waitErr)
			require.Equal(t, openapi1.FAILED, query.GetStatus(),
				"COPY to non-existent table should fail")
		} else {
			require.Equal(t, http.StatusBadRequest, resp.StatusCode,
				"COPY to non-existent table should return 400")
		}
	})

	t.Run("Copy_WithInvalidPath_Fails", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "people")

		queryId, resp, err := SubmitCopyQueryV1(dbClient, ctx,
			"/nonexistent/path/data.csv",
			"people",
			true)

		if err == nil && resp.StatusCode == http.StatusOK {
			query, waitErr := WaitForQueryCompletionV1(dbClient, ctx, queryId, 10*time.Second)
			require.NoError(t, waitErr)
			require.Equal(t, openapi1.FAILED, query.GetStatus(),
				"COPY with invalid path should fail")
		}
	})

	t.Run("Copy_InvalidPath_NoPartialData", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "people")

		// Try to COPY from invalid path
		queryId, resp, _ := SubmitCopyQueryV1(dbClient, ctx,
			"/nonexistent/path/data.csv",
			"people",
			true)

		if resp != nil && resp.StatusCode == http.StatusOK {
			WaitForQueryCompletionV1(dbClient, ctx, queryId, 10*time.Second)
		}

		// Table should remain empty (no partial data)
		result := ExecuteSelectStarV1(t, dbClient, ctx, "people")
		rows := ParseQueryResultsV1(result)
		require.Empty(t, rows, "Table should be empty after failed COPY")
	})

	t.Run("Copy_MultipleTimes_AccumulatesData", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "types_test")

		// First COPY
		LoadTestDataV1(t, dbClient, ctx, "types_test")
		result1 := ExecuteSelectStarV1(t, dbClient, ctx, "types_test")
		count1 := CountRowsV1(result1)
		require.Greater(t, count1, 0, "Should have rows after first COPY")

		// Second COPY - same data
		LoadTestDataV1(t, dbClient, ctx, "types_test")
		result2 := ExecuteSelectStarV1(t, dbClient, ctx, "types_test")
		count2 := CountRowsV1(result2)

		require.Equal(t, count1*2, count2,
			"Second COPY should double the row count")
	})

	t.Run("Copy_WithHeader_ParsesCorrectly", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "people")

		// COPY with header=true
		queryId, resp, err := SubmitCopyQueryV1(dbClient, ctx,
			"/data/tables/people/data.csv",
			"people",
			true)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		query, err := WaitForQueryCompletionV1(dbClient, ctx, queryId, 30*time.Second)
		require.NoError(t, err)
		require.Equal(t, openapi1.COMPLETED, query.GetStatus())

		// Verify correct number of rows (header not counted as data)
		result := ExecuteSelectStarV1(t, dbClient, ctx, "people")
		rows := ParseQueryResultsV1(result)
		require.Len(t, rows, 5, "Should have 5 data rows (header excluded)")
	})

	t.Run("Copy_EmptyTable_SelectReturnsEmpty", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "people")
		// No COPY executed

		result := ExecuteSelectStarV1(t, dbClient, ctx, "people")
		rows := ParseQueryResultsV1(result)
		require.Empty(t, rows, "Empty table should return no rows")
	})
}

// ============================================================================
// COPY Atomicity Tests (Interface v1)
// ============================================================================

func TestV1_Copy_Atomicity(t *testing.T) {
	RequireInterfaceVersion(t, 1)

	dbClient := DbClientV1(BaseURL)
	ctx := context.Background()

	t.Run("Copy_FailureDoesNotAffectExistingData", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "types_test")

		// First, load some valid data
		LoadTestDataV1(t, dbClient, ctx, "types_test")
		result1 := ExecuteSelectStarV1(t, dbClient, ctx, "types_test")
		count1 := CountRowsV1(result1)
		require.Greater(t, count1, 0, "Should have data after first COPY")

		// Try to COPY from invalid path
		queryId, resp, _ := SubmitCopyQueryV1(dbClient, ctx,
			"/nonexistent/invalid/path.csv",
			"types_test",
			true)

		if resp != nil && resp.StatusCode == http.StatusOK {
			WaitForQueryCompletionV1(dbClient, ctx, queryId, 10*time.Second)
		}

		// Original data should still be there
		result2 := ExecuteSelectStarV1(t, dbClient, ctx, "types_test")
		count2 := CountRowsV1(result2)
		require.Equal(t, count1, count2,
			"Failed COPY should not affect existing data")
	})

	t.Run("Copy_QueryStatusTransitions", func(t *testing.T) {
		_ = SetupTestTableV1(t, dbClient, ctx, "people")

		queryId, resp, err := SubmitCopyQueryV1(dbClient, ctx,
			"/data/tables/people/data.csv",
			"people",
			true)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		t.Log(pit.FormatResponse(resp))

		// Query should eventually complete
		query, err := WaitForQueryCompletionV1(dbClient, ctx, queryId, 30*time.Second)
		require.NoError(t, err)
		require.Equal(t, openapi1.COMPLETED, query.GetStatus())
	})
}
