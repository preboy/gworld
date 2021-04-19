@echo on

go build -o plan/exporter.exe       tools\exporter\main.go
go build -o plan/exporter_map.exe   tools\exporter_map\main.go

echo Done !!!

pause
