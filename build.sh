#!/bin/bash

# Build the application
GOEXPERIMENT=arenas go build -o arena-example src/*.go  && ./arena-example


