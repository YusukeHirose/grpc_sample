# grpc_sample

## 使用言語
go1.12.9 darwin/amd64

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

#### コード生成
```cd proto```
```protoc --go_out=plugins=grpc:./pb book.proto```

#### document生成
```go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc```
```protoc --doc_out=html,index.html:./ proto/*.proto```

## 実行手順
### grpccインストール
```npm install -g grpcc```

### grpcサーバー起動
```go run main.go```

### grpcc起動
```grpcc --proto ./proto/book.proto --address 127.0.0.1:8080 -i```

### リクエスト例
```client.getBook({title: "Docker"}, printReply)```


