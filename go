#!/usr/bin/env bash
set -aeuo pipefail

GOLANG_CONTAINER="$(docker ps -q -f NAME=golang)"
# docker exec -ti $WP_CLI sh -c "go $*"
docker-compose exec golang go $*
