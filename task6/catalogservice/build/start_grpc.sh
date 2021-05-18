#!/usr/bin/env bash

./wait-for-it.sh "${DB_HOST:-database}":"${DB_PORT:-3306}" -s -t 180 -- \
./main --mode grpc
