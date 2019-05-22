@echo off

node ./0.gen_list.js
node ./0.gen_files.js

call 0.build_go.bat
call 0.build_lua.bat

echo ALL IS OK !!!

pause