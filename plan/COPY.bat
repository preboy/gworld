@echo off


rmdir /S /Q ..\..\2dgame\simulator\win32\src\configuration
rmdir /S /Q ..\bin\configuration


mkdir ..\..\2dgame\simulator\win32\src\configuration
mkdir ..\bin\configuration


copy /Y lua\*.lua   ..\..\2dgame\simulator\win32\src\configuration\*.lua
copy /Y json\*.json ..\bin\configuration\*.json


echo COPY OVER !
pause