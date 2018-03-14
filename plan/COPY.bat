@echo off


rmdir /S /Q ..\..\2dgame\simulator\win32\src\configs_raw
rmdir /S /Q ..\bin\config


mkdir ..\..\2dgame\simulator\win32\src\configs_raw
mkdir ..\bin\config


copy /Y lua\*.lua   ..\..\2dgame\simulator\win32\src\configs_raw\*.lua
copy /Y json\*.json ..\bin\config\*.json


echo COPY OVER !
pause
