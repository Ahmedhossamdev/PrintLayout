#!/bin/bash

OUTPUT_DIR="bin"
mkdir -p $OUTPUT_DIR

PLATFORMS=(
  "windows/amd64"
  "linux/amd64"
  "darwin/amd64"
)

for PLATFORM in "${PLATFORMS[@]}"; do
  OS=$(echo $PLATFORM | cut -d'/' -f1)
  ARCH=$(echo $PLATFORM | cut -d'/' -f2)

  if [ "$OS" = "windows" ]; then
    OUTPUT_NAME="pr.exe"
  else
    OUTPUT_NAME="pr-$OS-$ARCH"
  fi

  echo "Building for $OS/$ARCH..."
  env GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUT_DIR/$OUTPUT_NAME" ./cmd/main.go
done

echo "Binaries built successfully in the $OUTPUT_DIR directory."
