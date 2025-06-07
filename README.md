# kubectllm
Kubectl with LLM

A command-line tool that converts natural language questions into kubectl commands using OpenAI.

## Usage

First, set your OpenAI API key:
```bash
export OPENAI_API_KEY="your-api-key-here"
```

### Ask Command
Generate kubectl commands for review:
```bash
kubectllm ask "list all the pods in the test namespace"
# Output: kubectl -n test get pods

kubectllm ask "get pod logs for my-app"  
# Output: kubectl logs my-app

kubectllm ask "describe deployment nginx"
# Output: kubectl describe deployment nginx
```

### Yolo Command
Generate and immediately execute kubectl commands:
```bash
kubectllm yolo "undeploy myapp"
# Output: kubectl delete deployment myapp
# ... followed by kubectl execution output ...

kubectllm yolo "list all pods in default namespace"
# Output: kubectl get pods
# ... followed by actual pod list ...
```

**WARNING:** The `yolo` command executes kubectl commands directly. Use with caution!

## Installation

```bash
go build -o kubectllm .
```

## Features

- Converts natural language to kubectl commands using OpenAI GPT-3.5-turbo
- **ask**: Prints generated commands for review before execution
- **yolo**: Generates and immediately executes kubectl commands
- Supports common kubectl operations
- Validates generated commands start with "kubectl" for security
- Requires OPENAI_API_KEY environment variable
