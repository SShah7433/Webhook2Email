#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o ../dist/webhook2email server.go