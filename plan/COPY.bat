@echo off


rmdir /S /Q ..\..\2dgame\simulator\win32\src\configs_json\quests
rmdir /S /Q ..\..\2dgame\simulator\win32\src\configs_raw
rmdir /S /Q ..\bin\config



mkdir ..\..\2dgame\simulator\win32\src\configs_json\quests
mkdir ..\..\2dgame\simulator\win32\src\configs_raw
mkdir ..\bin\config\quests


copy /Y lua\*.lua   ..\..\2dgame\simulator\win32\src\configs_raw\*.lua
copy /Y json\*.json ..\bin\config\*.json


copy /Y manual\quests\*.json ..\..\2dgame\simulator\win32\src\configs_json\quests\
copy /Y manual\quests\*.json ..\bin\config\quests\


echo COPY OVER !
pause
