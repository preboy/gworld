@echo off

call 0.build_go.bat
call 0.build_lua.bat

node ./0.gen_files.js

echo ALL IS OK !!!
pause