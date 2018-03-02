@echo off

forfiles /M *.proto /C "cmd /c protoc --gogofaster_out=..\msg @file"
echo Done
pause