# chatbot-DSL -- Backend

## gRPC generation

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/chat.proto
```

## Run

```bash
go run main.go
```

## Test

```bash
go test ./...
```
