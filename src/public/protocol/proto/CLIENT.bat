@echo off

protoc --descriptor_set_out ../protocol.pb session.proto game.proto

move ..\protocol.pb ..\..\..\..\..\2dgame\simulator\win32\src\message\protocol.pb

echo Done
pause