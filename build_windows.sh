#!/bin/bash

rm -rf ./build
mkdir ./build

CC=x86_64-w64-mingw32-gcc 
CGO_ENABLED=1 
GOOS=windows 
GOARCH=amd64
go build -v -o ./build/client.exe github.com/yeahyeahcore/HardwareMonitorNET/cmd/client 