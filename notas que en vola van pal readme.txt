comandos para compilar los protos en pb.go:
- protoc --go_out=plugins=grpc:clientelogistica clientelogistica/clientelogistica.proto
- protoc --go_out=plugins=grpc:camionlogistica camionlogistica/camionlogistica.proto


maquina de logistica debe contener carpetas:
- /logistica
- /camionlogistica
- /clientelogistica

se debe ejecutar desde la carpeta que contiene a logistica, clientelogistica, financiero,camionlogistica, camion, etc