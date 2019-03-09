@echo off

echo generate proto file for golang ...

protoc --gogofaster_out=..\msg *.proto

echo done