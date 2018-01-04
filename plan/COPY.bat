@echo off


rmdir /S /Q ..\..\2dgame\simulator\win32\src\config
rmdir /S /Q ..\bin\config


mkdir ..\..\2dgame\simulator\win32\src\config
mkdir ..\bin\config


copy /Y lua\*.lua   ..\..\2dgame\simulator\win32\src\config\*.lua
copy /Y json\*.json ..\bin\config\*.json


echo COPY OVER !
pause