#!/bin/sh

# -----------------------------------------------------------------------------

echo "clear ..."
go clean -cache

./ddz_build.sh
./game_build.sh
