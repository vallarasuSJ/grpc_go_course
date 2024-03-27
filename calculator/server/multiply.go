package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func (s *Server) Multiply(ctx context.Context,in *pb.MultiplyRequest) (*pb.MultiplyResponse, error){
	log.Printf("Multiply function was invoked with %v\n",in)
	return &pb.MultiplyResponse{
		Result: in.First_Number*
		in.Second_Number,
	},nil
}