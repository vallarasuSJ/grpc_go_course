package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func (s *Server) Subtract(ctx context.Context,in *pb.SubRequest) (*pb.SubResponse, error){
	log.Printf("Subtract function was invoked %v\n",in)
	
	firstNumber:=in.FirstNumber
	secondNumber:=in.SecondNumber

	var result int32
	if firstNumber>secondNumber{
		result=firstNumber-secondNumber
	}else{
		result=secondNumber-firstNumber
	}

	return &pb.SubResponse{
		Result: result,
	},nil

	

}