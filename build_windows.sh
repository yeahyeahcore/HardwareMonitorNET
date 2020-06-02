#!/bin/bash

rm -rf ./build
mkdir ./build

export CC=x86_64-w64-mingw32-gcc 
export CGO_ENABLED=1 
export GOOS=windows 
export GOARCH=amd64
go build -v -o ./build/client.exe github.com/yeahyeahcore/HardwareMonitorNET/cmd/client 