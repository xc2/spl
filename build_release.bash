#!/bin/bash

RELEASE_DIR="release"
BINARY_NAME="smptpl"

TARGETS="""
linux/amd64
linux/arm64
windows/amd64
windows/arm
darwin/amd64
darwin/arm64
"""

mkdir -p "$RELEASE_DIR"

echo "$TARGETS" | while IFS='/' read -r _os _arch; do

  if test -n "$_os" && test -n "$_arch"; then
    env CGO_ENABLED=0 GOOS="$_os" GOARCH="$_arch" go build -ldflags='-w -s' -buildmode=exe -o "${RELEASE_DIR}/${BINARY_NAME}-${_os}-${_arch}" .
  fi
done

# go build -ldflags='-w -s' -v -o "${RELEASE_DIR}/${BINARY_NAME}" .
