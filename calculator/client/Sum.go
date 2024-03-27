package main

import (
	"context"
	"log"
	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient){
	log.Println("doSum was invoked")
	res, err:=c.Sum(context.Background(),&pb.SumRequest{
		FirstNumber: 1,
		SecondNumber: 1,
	})

	if err!=nil{
		log.Fatalf("Could not sum: %v",err)
	}
	log.Printf("Sum %d\n",res.Result)


}