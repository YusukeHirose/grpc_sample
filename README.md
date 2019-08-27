# grpc_sample

## 使用言語
go1.12.6 darwin/amd64

## 環境構築
### go module使用準備
```export GO111MODULE=on```
```go mod init grpc_sample```

### gRPCインストール
```go get -u google.golang.org/grpc```

### protocインストール
```brew install protobuf```

### protocプラグインインストール
```go get -u github.com/golang/protobuf/protoc-gen-go```

