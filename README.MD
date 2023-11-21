# At first install `protoc` and add to PATH variable 

## Compile proto to GO

1. install plugins:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
2. add some option to proto `option go_package = "illinois/grcp/echo";`
3. build proto:
```sh
mkdir pkg
protoc --go_out=./pkg/ --go_opt=paths=source_relative --go-grpc_out=./pkg/ --go-grpc_opt=paths=source_relative file.proto
```

## Compile proto to Python
1. install plugins:
```sh
python -m pip install grpcio
python -m pip install grpcio-tools
```
2. build proto:
```sh
mkdir pkg
python -m grpc_tools.protoc -I./ --python_out=./pkg/ --grpc_python_out=./pkg/ file.proto
```