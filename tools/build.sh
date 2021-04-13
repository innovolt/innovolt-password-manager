#!/bin/bash

set -x

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

rm -rf $SCRIPTPATH/artifacts
mkdir $SCRIPTPATH/artifacts


ROOTPATH="$(dirname "$SCRIPTPATH")"

pushd $ROOTPATH
# Build go executable binary
GOOS=linux GOARCH=amd64 go build
cp innovolt-pm $SCRIPTPATH/artifacts/
popd

# build docker image
docker build -t innovolt-pm -f $SCRIPTPATH/Dockerfile $SCRIPTPATH
