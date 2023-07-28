#!/bin/bash

# Run all tests
# Usage: ./run-tests.sh
#
GOEXPERIMENT=arenas go test -v ./... -bench=BenchmarkGarbageCollector -benchtime=5m -benchmem
GOEXPERIMENT=arenas go test -v ./... -bench=BenchmarkArenas -benchtime=5m -benchmem

#GOEXPERIMENT=arenas go test -v ./... -bench=. -benchtime=10s -benchmem

