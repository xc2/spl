#!/bin/bash

RELEASE_DIR="${RELEASE_DIR:-release}"
BINARY_NAME="${BINARY_NAME:-spl}"
RELEASE_VER="${RELEASE_VER:-}"

fn="${BINARY_NAME}"
if test -n "$RELEASE_VER"; then
  fn="${fn}-${RELEASE_VER}"
fi

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
    env CGO_ENABLED=0 GOOS="$_os" GOARCH="$_arch" go build -ldflags="-w -s -X main.CliName=${BINARY_NAME} -X main.CliVersion=${RELEASE_VER}" -buildmode=exe -o "${RELEASE_DIR}/${fn}-${_os}-${_arch}" .
  fi
done
