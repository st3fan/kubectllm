package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestAskCommandRequiresAPIKey(t *testing.T) {
	// Ensure OPENAI_API_KEY is not set
	originalKey := os.Getenv("OPENAI_API_KEY")
	os.Unsetenv("OPENAI_API_KEY")
	defer func() {
		if originalKey != "" {
			os.Setenv("OPENAI_API_KEY", originalKey)
		}
	}()

	// Test that ask command requires OPENAI_API_KEY
	err := runAskCommand(nil, []string{"test question"})
	if err == nil {
		t.Error("Expected error when OPENAI_API_KEY not set, but got none")
	}
	if !strings.Contains(err.Error(), "OPENAI_API_KEY environment variable is required") {
		t.Errorf("Expected error about OPENAI_API_KEY, got: %v", err)
	}
}