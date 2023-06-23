# config-service

Проект содержит серверное и клиентское приложения, которые позволяют настраивать hostname и список DNS-серверов на сервере с клиента. Сервер представляет gRPC и REST API, а также генерирует документацию по OpenAPI
спецификации. Она будет доступна по адресу <your url>/swagger/, а также локально в директории api/grpcGen/config.swagger.json

## Зависимости
- go
- protobuf
- github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
- github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
- google.golang.org/protobuf/cmd/protoc-gen-go
- google.golang.org/grpc/cmd/protoc-gen-go-grpc

Установка для Arch Linux
```bash
sudo pacman -S go protobuf
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
  
## Сборка

```bash
cd build
make all
```

## Использование
```bash
cd cmd/bin
# сервер
./service -G <your grpc port> -H <your http gateway port> -C <your connection name in networkmanager>
# клиент
./client hostname -d <grpc url>
```
Для более подробной информации по использованию запустите бинарник со флагов --help
