package main

import (
	"context"
	"log"

	pb "github.com/ararage/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum Function was invoked %v\n", in)
	result := in.FirstNumber + in.SecondNumber
	return &pb.SumResponse{
		Result: result,
	}, nil
}
