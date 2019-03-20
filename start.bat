@echo off

del bin\*.exe

set GOPATH=%cd%
go install server

if %ERRORLEVEL% NEQ 0 goto ERROR

cd bin
server.exe
cd ..
goto END

:ERROR
echo "ERROR on INSTALL"

:END
