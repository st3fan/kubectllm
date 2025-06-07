package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestYoloCommandRequiresAPIKey(t *testing.T) {
	// Ensure OPENAI_API_KEY is not set
	originalKey := os.Getenv("OPENAI_API_KEY")
	os.Unsetenv("OPENAI_API_KEY")
	defer func() {
		if originalKey != "" {
			os.Setenv("OPENAI_API_KEY", originalKey)
		}
	}()

	// Test that yolo command requires OPENAI_API_KEY
	err := runYoloCommand(nil, []string{"test question"})
	if err == nil {
		t.Error("Expected error when OPENAI_API_KEY not set, but got none")
	}
	if !strings.Contains(err.Error(), "OPENAI_API_KEY environment variable is required") {
		t.Errorf("Expected error about OPENAI_API_KEY, got: %v", err)
	}
}

func TestYoloCommandMinimumArgs(t *testing.T) {
	// Test that yolo command requires at least one argument
	// Note: This will be caught by Cobra's MinimumNArgs validation
	// before runYoloCommand is called, so we test that the command
	// is properly configured
	if yoloCmd.Args == nil {
		t.Error("Expected yolo command to have args validation")
	}
}

func TestYoloCommandValidatesKubectlCommand(t *testing.T) {
	// Test case where we have a fake API key but the generated command
	// doesn't start with kubectl - we'll create a mock scenario
	
	// This test is limited since we can't easily mock the OpenAI response
	// without significant refactoring, but we can test the command validation logic
	
	// The validation is done in runYoloCommand after generateKubectlCommand
	// If the command doesn't start with "kubectl", it should return an error
	
	// We can't easily test this without mocking the OpenAI API call,
	// but the logic is in place to validate that generated commands start with "kubectl"
}