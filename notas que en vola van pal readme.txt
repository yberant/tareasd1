comandos para compilar los protos en pb.go:
- protoc --go_out=plugins=grpc:clientelogistica clientelogistica/clientelogistica.proto
- protoc --go_out=plugins=grpc:camionlogistica camionlogistica/camionlogistica.proto
