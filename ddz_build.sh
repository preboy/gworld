#!/bin/sh


# -----------------------------------------------------------------------------
platform=`go env GOOS`
echo "Installing DDZ for '$platform' ..."

if [ $platform = "windows" ];
then
    go build -o bin_ddz/ddz.exe     -gcflags="-N -l" ddz/main.go
    go build -o bin_ddz/ddz_ai.exe  -gcflags="-N -l" ddz_ai/main.go
    go build -o bin_ddz/ddz_rr.exe  -gcflags="-N -l" ddz_rr/main.go
else
    go build -o bin_ddz/ddz         -gcflags="-N -l" ddz/main.go
    go build -o bin_ddz/ddz_ai      -gcflags="-N -l" ddz_ai/main.go
    go build -o bin_ddz/ddz_rr      -gcflags="-N -l" ddz_rr/main.go
fi;


# -----------------------------------------------------------------------------
echo -e "\033[32mBuild Done\033[0m"
