#!/bin/bash

# Check if go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Run the Go program
go run rewards.go

# Check if the Go program executed successfully
if [ $? -ne 0 ]; then
    echo "Error: Failed to run rewards.go"
    exit 1
else
    echo "rewards.go executed successfully."
fi

