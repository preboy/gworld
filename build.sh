#!/bin/sh

go env -w GO111MODULE=off
export GOPATH=`pwd`
echo GOPATH=$GOPATH

# -----------------------------------------------------------------------------

echo "Installing server ..."
go install -gcflags="-N -l" game
go install -gcflags="-N -l" router


# -----------------------------------------------------------------------------

echo -e "\033[32mBuild Done\033[0m"
