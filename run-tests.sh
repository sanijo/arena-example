#!/bin/bash

# Run all tests
# Usage: ./run-tests.sh
#
GOEXPERIMENT=arenas go test -v ./... -bench=. -benchtime=10s

