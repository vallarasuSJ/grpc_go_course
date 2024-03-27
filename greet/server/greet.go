package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/greet/proto"
)

func (s *Server) Greet (ctx context.Context,in *pb.GreetRequest) (*pb.GreetResponse, error){
	log.Printf("Greet function was invoked with %v\n",in)

	return &pb.GreetResponse{
		Result: "Hello "+ in.FirstName,
	},nil
}