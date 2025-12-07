package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/smogork/ISBD-MIMUW/pit"
)

var BaseURL string

func TestMain(m *testing.M) {
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
