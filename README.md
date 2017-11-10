
## prepare third package 
```
go get gopkg.in/mgo.v2
go get github.com/gogo/protobuf/gogoproto
go get github.com/gogo/protobuf/protoc-gen-gofast
go get github.com/golang/protobuf/proto
go generate
```
protoc  --gofast_out=. login.proto
