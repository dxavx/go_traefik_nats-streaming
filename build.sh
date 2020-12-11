#!/usr/bin/env bash

docker-compose -f ./docker-compose.yml up -d --scale subscriber=1
#docker-compose -f ./docker-compose.yml up -d
