@echo off

set GOPATH=%cd%
go install server

cd bin
server.exe
