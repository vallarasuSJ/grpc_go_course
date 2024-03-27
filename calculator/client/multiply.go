package main

import (
	"context"
	"log"
	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func doMultiply(c pb.CalculatorServiceClient){
	log.Println("doMultiply was invoked")
	res,err:=c.Multiply(context.Background(),&pb.MultiplyRequest{
		First_Number: 7,
		Second_Number: 6,
	})

	if err!=nil{
		log.Fatalf("Could not sum: %v",err)
	}

	log.Printf("Sum: %d\n",res.Result)
}