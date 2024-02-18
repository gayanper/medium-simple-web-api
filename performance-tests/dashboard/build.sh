#!/bin/sh

docker run --rm -it -e GOOS=darwin -v "${PWD}:/xk6" \
  grafana/xk6 build v0.43.1 \
  --with github.com/grafana/xk6-dashboard@latest