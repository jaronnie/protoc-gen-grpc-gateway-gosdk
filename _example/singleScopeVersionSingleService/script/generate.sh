#!/usr/bin/env bash

set -e

Mode=$1

function generateFunc() {
    mkdir -p autosdk/pb
    for dir in proto/*; do
      dir_name="${dir##*/}"
      if [ -d "$dir" ]; then
          protoc --proto_path="$dir" --go_out="$(pwd)"/autosdk/pb --grpc-gateway-gosdk_out=logtostderr=true,v=1,env_file=./conf/cfg.toml:"$(pwd)"/autosdk "$dir"/*.proto
      fi
    done
    sleep 1
    cd autosdk; go mod tidy; cd ..
}

generateFunc "$Mode"