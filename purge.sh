#!/usr/bin/env bash

docker-compose down --remove-orphans
docker image rm go_traefik_nats_api
docker image rm go_traefik_nats_subscriber
