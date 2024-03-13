#!/bin/bash

# Build Linux binary
go build -o monkey ./src/main.go

# Build Windows binary
GOOS=windows go build -o monkey.exe ./src/main.go
