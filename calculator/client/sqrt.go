package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func doSqrt(c pb.CalculatorServiceClient){
	log.Println("Sqrt was invoked")  
	var n int32;
	log.Println("Enter the number: ")
	fmt.Scanln(&n)
	res, err:=c.Sqrt(context.Background(),&pb.SqrtRequest{
		Number:n,
	}) 

	if err!=nil{
		e,ok:=status.FromError(err) 

		if ok{
			log.Printf("Error message from server: %s\n",e.Message())
			log.Printf("Error message from server: %s\n",e.Code()) 

			if e.Code()==codes.InvalidArgument{
				log.Println("We probably sent a negative number!")
			}

		}else{
			log.Fatalf("A non gRPC error: %v\n",err)
		}

	}

	log.Printf("Sqrt: %f\n",res.Result)
}