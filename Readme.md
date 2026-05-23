# Installing GRPC and protobuffer plugins for golang
Protobuffers :
```
go install gooogle.golang.org/protobuf/cmd/protoc-gen-go@latest
```
```
GRPC :  `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
Deoendencies :
go get google.golang.org/protobuf
go get google.golang.org/grpc
go get google.golang.org/genproto
Installing protobuf compiler for linux (protoc compiler)
sudo apt install -y protobuf-compiler
or 
sudo pacman -S protobuf

