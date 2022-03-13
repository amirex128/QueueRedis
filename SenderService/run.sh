#! /bin/bash

# run only mysql and redis server from docker-compose.yml
docker-compose up -d mysql redis-server
go run main.go