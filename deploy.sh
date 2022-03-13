#! /bin/bash

# run all services in the docker-compose file (SenderService, ReceiverService, Redis, Mysql, PhpMyAdmin) for production
docker-compose up -d mysql redis-server
