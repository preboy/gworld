@echo off
protoc  --gogofaster_out=..\msg *.proto
echo Done
pause