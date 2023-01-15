#!/bin/bash
set -e

# Set the build variables and export them
APP_NAME=${APP_NAME:-"zwitter-server"}
VERSION=$(cat ../VERSION)
BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT=$(git rev-parse HEAD)
BUILD_NO=${BUILD_NO:-0}

# Create the version.json file
cat <<EOF > ../cmd/${APP_NAME}/version.txt
${VERSION}
${BUILD_NO}
${BUILD_DATE}
${GIT_COMMIT}
EOF


# Build the binary
cd ..
go build ./cmd/zwitter-server