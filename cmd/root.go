package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubectllm",
	Short: "Kubectl with LLM - Convert natural language to kubectl commands",
	Long: `kubectllm uses OpenAI to convert natural language questions into kubectl commands.
The generated commands are printed to the console for review before execution.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands here
	rootCmd.AddCommand(askCmd)
}