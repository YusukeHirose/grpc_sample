package main

import (
	pb "grpc_sample/pb"
	"grpc_sample/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	bookService := &service.BookService{}
	log.Println("started!!")
	// 実行したい実処理をseverに登録する
	pb.RegisterBookServer(server, bookService)
	server.Serve(listenPort)
}
