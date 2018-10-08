#!/bin/sh

export GOPATH=$(pwd)
echo GOPATH=$GOPATH

# -----------------------------------------------------------------------------

echo "Installing server ..."
go install $1  -gcflags="-N -l" server


# -----------------------------------------------------------------------------

echo -e "\033[32mDone\033[0m"