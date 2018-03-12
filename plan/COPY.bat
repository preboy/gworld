@echo off


rmdir /S /Q ..\..\2dgame\simulator\win32\src\configs
rmdir /S /Q ..\bin\config


mkdir ..\..\2dgame\simulator\win32\src\configs
mkdir ..\bin\config


copy /Y lua\*.lua   ..\..\2dgame\simulator\win32\src\configs\*.lua
copy /Y json\*.json ..\bin\config\*.json


echo COPY OVER !
pause
