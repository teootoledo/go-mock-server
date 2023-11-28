#!/bin/bash

docker_buildkit=0 docker build -f Dockerfile -t api-service .

docker-compose up -d
