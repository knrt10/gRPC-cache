#!/bin/bash
go build -o server-cache api/main.go && go build -o client-cache examples/client.go
