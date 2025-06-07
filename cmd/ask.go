package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask [question]",
	Short: "Ask a natural language question to generate a kubectl command",
	Long: `Convert a natural language question into a kubectl command using OpenAI.
The generated command is printed to the console for review.

Example:
  kubectllm ask "list all the pods in the test namespace"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runAskCommand,
}

func runAskCommand(cmd *cobra.Command, args []string) error {
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
	return nil
}

func generateKubectlCommand(question, apiKey string) (string, error) {
	client := openai.NewClient(apiKey)

	prompt := fmt.Sprintf(`Convert the following natural language question into a kubectl command. Return only the kubectl command, nothing else.

Question: %s

Examples:
- "list all pods" → "kubectl get pods"
- "list all pods in the test namespace" → "kubectl -n test get pods"
- "get pod logs for my-pod" → "kubectl logs my-pod"
- "describe pod my-pod" → "kubectl describe pod my-pod"
- "delete pod my-pod in namespace test" → "kubectl -n test delete pod my-pod"

kubectl command:`, question)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens:   100,
			Temperature: 0.1, // Low temperature for more consistent output
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	command := strings.TrimSpace(resp.Choices[0].Message.Content)
	return command, nil
}