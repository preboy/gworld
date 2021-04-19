@echo off

del bin\game.exe

rem set GOPATH=%cd%
go build -o bin/game.exe  -gcflags="-N -l" game/main.go

if %ERRORLEVEL% NEQ 0 goto ERROR

cd bin
game.exe
cd ..
goto END

:ERROR
echo "ERROR on INSTALL"

:END
