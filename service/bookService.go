package service

import (
	"context"
	"errors"

	pb "grpc_sample/pb"
)

type BookService struct{}

func (b *BookService) GetBook(ctx context.Context, title *pb.GetBookTitle) (*pb.BookResponse, error) {
	switch title.Title {
	case "Docker":
		return &pb.BookResponse{
			Title:  "Docker",
			Author: "author",
		}, nil
	case "Web API ~":
		return &pb.BookResponse{
			Title:  "Web API ~",
			Author: "abc",
		}, nil
	}
	return nil, errors.New("Not Found YourCat")
}
