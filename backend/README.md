# chatbot-DSL -- Backend

## gRPC

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/chat.proto
```