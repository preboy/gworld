@echo off

protoc --descriptor_set_out ../protocol.pb session.proto game.proto

move ..\protocol.pb ..\..\..\..\..\2dgame\simulator\win32\protocol.pb

echo Done
pause