@echo off


del /S /Q ..\..\whome\simulator\win32\src\config\quests\*
del /S /Q ..\..\whome\simulator\win32\src\config\excel\*


del /S /Q ..\bin\config\quests\*
del /S /Q ..\bin\config\excel\*


copy /Y lua\*.lua   ..\..\whome\simulator\win32\src\config\excel\*.lua
copy /Y json\*.json ..\bin\config\excel\*.json


copy /Y manual\quests\*.json ..\..\whome\simulator\win32\src\config\quests\*.json
copy /Y manual\quests\*.json ..\bin\config\quests\*.json


echo COPY OVER !
pause
