package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func doFibonacci(c pb.CalculatorServiceClient){
	log.Println("doFibonacci function was invoked")

	var number int32;

	fmt.Println("Enter the fibonacci number: ")
	if _,err:=fmt.Scanln(&number); err!=nil{
		log.Fatalf("Error")
	}  

	req:=&pb.FibRequest{
		FibNumber: number,
	}
	res,err:=c.Fibonacci(context.Background(),req)

	if err!=nil{
		log.Fatalf("Error while opening the connection %v\n",err)
	}
	for{
		result,err:=res.Recv()

		if err==io.EOF{
			break
		}

		if err!=nil{
			log.Fatalf("Error while receiving data from the server")
		}

		log.Printf("Fibonacci number: %v\n",result.Result)
		time.Sleep(1*time.Second)

	}




}