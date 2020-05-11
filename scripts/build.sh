#!/bin/bash
go build -o server-cache pkg/server/main.go && go build -o client-cache pkg/client/client.go
