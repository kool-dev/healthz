#!/bin/bash

while [ $# -gt 0 ]; do
  case "$1" in
    --version )
      BUILD_VERSION="$2"
      shift 2
      ;;
    -- )
      shift
      break
      ;;
    *)
      echo "Invalid Argument: ${1}"
      exit 1
      ;;
  esac
done

if [ -f .env ]; then
    source .env
fi

GO_IMAGE=${GO_IMAGE:-golang:1.15.0}

if [ "$BUILD_VERSION" == "" ]; then
    echo "missing environment variable BUILD_VERSION"
    exit 5
fi

read -p "You are going to build all artifacts for version $BUILD_VERSION. Continue? (y/N) "
if [[ ! $REPLY =~ ^(yes|YES|y|Y)$ ]]
then
   exit
fi

rm -rf dist
mkdir -p dist

# ATTENTION - binary names must match the -GOOS-GOARCH suffix
# because self-update relies on this pattern to work.
BUILD=(\
  "dist/healthz-linux-amd64|--env GOOS=linux --env GOARCH=amd64" \
  "dist/healthz-linux-arm6|--env GOOS=linux --env GOARCH=arm --env GOARM=6" \
  "dist/healthz-linux-arm7|--env GOOS=linux --env GOARCH=arm --env GOARM=7" \
)

for i in "${!BUILD[@]}"; do
    dist=$(echo ${BUILD[$i]} | cut -d'|' -f1)
    flags=$(echo ${BUILD[$i]} | cut -d'|' -f2)
    echo "Building to ${flags}"
    docker run --rm \
        $flags \
        --env CGO_ENABLED=0 \
        -v $(pwd):/code -w /code $GO_IMAGE \
        go build -a -tags 'osusergo netgo static_build' \
        -ldflags '-X kool-dev/healthz/cmd.version='$BUILD_VERSION' -extldflags "-static"' \
        -o $dist
done

echo "Going to generate CHECKSUMS"

for file in dist/*; do
    shasum -a 256 $file > $file.sha256
done

echo "Finished building all artifacts for version $BUILD_VERSION"
