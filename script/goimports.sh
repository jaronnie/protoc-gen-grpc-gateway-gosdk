#!/bin/bash

set -e

WorkPath=$(pwd)

FmtType=$1

# export tools bin to PATH
PATH="${WorkPath}"/tools:"${PATH}"

if [ ! -f "$WorkPath"/tools/goimports-reviser-darwin ]; then
  {
    mkdir -p tools
    wget -P "${WorkPath}"/tools https://nexus.hyperchain.cn/repository/blocface/blocface/tools/goimports-reviser-darwin
    chmod +x "${WorkPath}"/tools/goimports-reviser-darwin
  }
fi

if [ "$FmtType" = "all" ]; then
  for i in $(find . -name "*.go" | grep -v ".pb.go" | grep -v ".pb.gw.go") ; do
    "$WorkPath"/tools/goimports-reviser-darwin -rm-unused -set-alias -format -file-path "$i"
  done
else
  for i in $(git diff --cached --name-only --diff-filter=ACM -- '*.go' | grep -v ".pb.go" | grep -v ".pb.gw.go") ; do
      "$WorkPath"/tools/goimports-reviser-darwin -rm-unused -set-alias -format -file-path "$i"
      git add "$i"
  done
fi