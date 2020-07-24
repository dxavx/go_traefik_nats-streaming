#!/usr/bin/env bash

docker-compose -f ./docker-compose.yml up -d --scale api=2
#docker-compose -f ./docker-compose.yml up -d
