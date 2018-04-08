@echo on

set GOPATH=%cd%
go install tools\exporter
go install tools\exporter_map

move /y bin\exporter.exe plan\
move /y bin\exporter_map.exe plan\

echo Done !!!
pause
