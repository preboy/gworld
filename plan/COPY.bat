@echo off

copy /Y lua\*.lua   ..\..\2dgame\simulator\win32\src\configuration\*.lua
copy /Y json\*.json ..\bin\configuration\*.json


echo COPY OVER !
pause