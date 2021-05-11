@echo off

del *.pb.go

protoc --gogofaster_out=. gambler.proto
protoc --gogofaster_out=. referee.proto

pause
