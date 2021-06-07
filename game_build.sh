#!/bin/sh


# -----------------------------------------------------------------------------
platform=`go env GOOS`
echo "Installing server for '$platform' ..."

if [ $platform = "windows" ];
then
    go build -o bin/game.exe  -gcflags="-N -l" game/main.go
else
    go build -o bin/game  -gcflags="-N -l" game/main.go
fi;


# -----------------------------------------------------------------------------
echo -e "\033[32mBuild Done\033[0m"
