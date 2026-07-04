package alloydb_test

import (
	"os"
	"testing"

	"github.com/whitehathub-007/langchaingo/internal/testutil/testctr"
)

func TestMain(m *testing.M) {
	testctr.EnsureTestEnv()
	os.Exit(m.Run())
}
