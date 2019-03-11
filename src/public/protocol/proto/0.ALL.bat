@echo off

node ./0.gen_list.js

call 0.build_go.bat
call 0.build_lua.bat

echo ALL IS OK !!!

pause