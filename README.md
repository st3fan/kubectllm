# kubectllm
Kubectl with LLM

A command-line tool that converts natural language questions into kubectl commands using OpenAI.

## Usage

First, set your OpenAI API key:
```bash
export OPENAI_API_KEY="your-api-key-here"
```

Then ask natural language questions:
```bash
kubectllm ask "list all the pods in the test namespace"
# Output: kubectl -n test get pods

kubectllm ask "get pod logs for my-app"  
# Output: kubectl logs my-app

kubectllm ask "describe deployment nginx"
# Output: kubectl describe deployment nginx
```

## Installation

```bash
go build -o kubectllm .
```

## Features

- Converts natural language to kubectl commands using OpenAI GPT-3.5-turbo
- Prints generated commands for review before execution
- Supports common kubectl operations
- Requires OPENAI_API_KEY environment variable
