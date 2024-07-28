#!/bin/bash

echo "Running tests..."

# Set up test environment variables, if any
export TEST_ENV_VAR=value

# Run tests
go test ./...

echo "Tests completed."
