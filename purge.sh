#!/usr/bin/env bash

docker-compose down --remove-orphans
docker image rm go_traefik_nats-api
docker image rm go_traefik_nats-subscriber
