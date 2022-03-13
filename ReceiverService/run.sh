#! /bin/bash

docker-compose up -d mysql redis-server
go run main.go