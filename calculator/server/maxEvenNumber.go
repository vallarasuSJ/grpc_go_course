package main

import (
	"io"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
) 

func (s *Server) MaxEvenNumber(in pb.CalculatorService_MaxEvenNumberServer) error{
	log.Println("MaxEvenNumber function was invoked")

	var result int32=0

	for{
		res,err:=in.Recv() 

		if err==io.EOF{
			return in.SendAndClose(&pb.EvenResponse{
				Result: result,
			})
		}
		
		if err!=nil{
			log.Fatalf("Error while receiving value from client: %v",err)
		} 

		log.Println("Received value: ",res.EvenNumber) 

		if res.EvenNumber%2==0 && res.EvenNumber>result{
			result=res.EvenNumber
		}
	}
}