package main

import (
	"log"


	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
) 

func (s *Server) Fibonacci(in *pb.FibRequest,stream pb.CalculatorService_FibonacciServer) error{

	log.Println("Fibonacci function was invoked from the server") 

	var number int32=in.FibNumber
	var a ,b int32=0,1
	for i:=int32(0);i<number;i++{
		stream.Send(&pb.FibResponse{
			Result: a,
		})
		a,b=a+b,a
	}
	return nil

}