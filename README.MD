## Compile proto to GO

1. install protoc and add to PATH variable
2. install plagins:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
3. add some option to proto `option go_package = "illinois/grcp/echo";`
4. build proto
```sh
protoc --go_out=./pkg/ --go_opt=paths=source_relative --go-grpc_out=./pkg/ --go-grpc_opt=paths=source_relative file.proto
```