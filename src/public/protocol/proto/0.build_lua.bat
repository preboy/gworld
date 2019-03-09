@echo off

echo generate proto file for golang ...

node ./0.gen_list.js

set /p fs=<FILES

protoc --descriptor_set_out ../protocol.pb %fs%

move ..\protocol.pb ..\..\..\..\..\2dgame\simulator\win32\src\message\protocol.pb

del FILES

echo Done