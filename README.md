# CMD

```bash
# install the proto compile for golang
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# install the proto compile for grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# compile protofile and generate code in golang
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <dir>/<file.proto>
```

