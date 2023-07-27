#!/bin/bash

# Run all tests
# Usage: ./run-tests.sh
#
GOEXPERIMENT=arenas go test -v ./... -bench=BenchmarkGarbageCollector -benchtime=30s -benchmem
GOEXPERIMENT=arenas go test -v ./... -bench=BenchmarkArenas -benchtime=30s -benchmem

#GOEXPERIMENT=arenas go test -v ./... -bench=. -benchtime=10s -benchmem

