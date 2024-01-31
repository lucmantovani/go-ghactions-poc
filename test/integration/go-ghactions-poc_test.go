package integration

import (
	"testing"
)

func TestGoGHActionsPOCIntegrationTest(t *testing.T) {
	t.Run("TestGoGHActionsPOCIntegrationTest", func(t *testing.T) {
		t.Fatal("Fail integration test in CI")
	})
}
