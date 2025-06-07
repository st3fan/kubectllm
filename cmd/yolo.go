package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var yoloCmd = &cobra.Command{
	Use:   "yolo [question]",
	Short: "Ask a natural language question and immediately execute the kubectl command",
	Long: `Convert a natural language question into a kubectl command using OpenAI
and immediately execute it. The generated command is printed first, then executed.

Example:
  kubectllm yolo "list all the pods in the test namespace"
  
WARNING: This will execute kubectl commands directly. Use with caution!`,
	Args: cobra.MinimumNArgs(1),
	RunE: runYoloCommand,
}

func runYoloCommand(cmd *cobra.Command, args []string) error {
	question := strings.Join(args, " ")
	
	// Check for OpenAI API key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}

	// Generate kubectl command using OpenAI
	kubectlCmd, err := generateKubectlCommand(question, apiKey)
	if err != nil {
		return fmt.Errorf("failed to generate kubectl command: %w", err)
	}

	// Print the generated command
	fmt.Println(kubectlCmd)

	// Parse the command to extract kubectl and its arguments
	cmdParts := strings.Fields(kubectlCmd)
	if len(cmdParts) == 0 {
		return fmt.Errorf("generated command is empty")
	}

	// Ensure the command starts with kubectl
	if cmdParts[0] != "kubectl" {
		return fmt.Errorf("generated command does not start with kubectl: %s", kubectlCmd)
	}

	// Execute the kubectl command
	var kubectlArgs []string
	if len(cmdParts) > 1 {
		kubectlArgs = cmdParts[1:]
	}

	execCmd := exec.Command("kubectl", kubectlArgs...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("failed to execute kubectl command: %w", err)
	}

	return nil
}