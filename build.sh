#!/bin/bash

# ./go-bindata -o common/bindata.go ./assets

# And then don't forget to change the package within the file
read -p "Don't forget to change the package within the file"

mkdir -p bin/
GOOS=darwin GOARCH=386 go build -o bin/darwin-flash-encrypt
GOOS=windows GOARCH=386 go build -o bin/windows-flash-encrypt.exe
GOOS=linux GOARCH=arm go build -o bin/arm-flash-encrypt