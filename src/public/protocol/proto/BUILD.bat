@echo off

if "%1"=="" goto EXIT
echo BUILDING %~nx1

protoc  --gogofaster_out=..\msg %~nx1
echo Done

:EXIT
pause
