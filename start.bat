@echo off

del bin\*.exe

set GOPATH=%cd%
go install game

if %ERRORLEVEL% NEQ 0 goto ERROR

cd bin
game.exe
cd ..
goto END

:ERROR
echo "ERROR on INSTALL"

:END
