@echo off


taskkill /f /im ddz.exe
taskkill /f /im ddz_ai.exe
taskkill /f /im ddz_rr.exe

timeout 1
start cmd /k ddz.exe

timeout 1
start cmd /k ddz_rr.exe

timeout 1
start ai.bat name_1 dev_test

timeout 1
start ai.bat name_2 dev_test

timeout 1
start ai.bat name_3 dev_test
