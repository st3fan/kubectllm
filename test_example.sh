#!/bin/bash

# Example usage of kubectllm
# This script demonstrates how the tool should be used

echo "Building kubectllm..."
go build -o kubectllm .

echo ""
echo "Testing basic help functionality:"
./kubectllm --help

echo ""
echo "Testing ask command help:"
./kubectllm ask --help

echo ""
echo "Testing yolo command help:"
./kubectllm yolo --help

echo ""
echo "Testing without OPENAI_API_KEY (should show error):"
./kubectllm ask "list all pods"

echo ""
echo "Testing yolo without OPENAI_API_KEY (should show error):"
./kubectllm yolo "list all pods"

echo ""
echo "To test with real OpenAI API, set OPENAI_API_KEY and run:"
echo 'export OPENAI_API_KEY="your-api-key-here"'
echo './kubectllm ask "list all the pods in the test namespace"'
echo ""
echo "Expected output: kubectl -n test get pods"
echo ""
echo 'For immediate execution, use:'
echo './kubectllm yolo "undeploy myapp"'
echo "Expected output: kubectl delete deployment myapp"
echo "... followed by kubectl execution output ..."

# Clean up
rm -f kubectllm