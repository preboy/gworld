@echo off

node ./0.gen_file.js

set /p fs=<FILES

protoc --descriptor_set_out ../protocol.pb %fs%

move ..\protocol.pb ..\..\..\..\..\2dgame\simulator\win32\src\message\protocol.pb

echo Done
pause
