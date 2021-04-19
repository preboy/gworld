@echo off

echo generate proto file for golang ...

set /p fs=<FILES

protoc --gogofaster_out=..\msg %fs%

echo done