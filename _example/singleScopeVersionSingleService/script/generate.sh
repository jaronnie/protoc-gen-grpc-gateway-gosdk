#!/usr/bin/env bash

set -e

Mode=$1

function generateFunc() {
    mkdir -p autosdk/pb
    for dir in proto/*; do
      dir_name="${dir##*/}"
      if [ -d "$dir" ]; then
          if [ "$1" = "debug" ]; then
            protoc --proto_path="$dir" --go_out="$(pwd)"/autosdk/pb --go-httpsdk_out=logtostderr=true,v=1,pluginOutputPath=autosdk,scopeVersion="$dir_name",gatewayPrefix=/gateway,env_file=./conf/cfg.toml,debug=true:"$(pwd)"/autosdk "$dir"/*.proto
          else
            protoc --proto_path="$dir" --go_out="$(pwd)"/autosdk/pb --go-httpsdk_out=logtostderr=true,pluginOutputPath=autosdk,scopeVersion="$dir_name",gatewayPrefix=/gateway,env_file=./conf/cfg.toml:"$(pwd)"/autosdk "$dir"/*.proto
          fi
      fi
    done
    sleep 1
    cd autosdk; go mod tidy; cd ..
}

generateFunc "$Mode"