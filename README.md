gworld = golang world

## mgo
```
go get gopkg.in/mgo.v2
```

## protobuf
```
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/protoc-gen-gogofaster
go get github.com/gogo/protobuf/gogoproto
```
protoc  --gogofaster_out=. login.proto

## for tools\exporter
go get  github.com/tealeg/xlsx


项目介绍：
	基于golang实现的手游服务端框架，充分利用了golang goroutine的特性。

	数据库： mongodb

配套demo:
	whome