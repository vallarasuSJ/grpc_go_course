package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
) 

func doPrimes(c pb.CalculatorServiceClient){
	log.Println("doPrimes was invoked") 

	req:=&pb.PrimeRequest{
		Number: 4324324,
	}

	stream,err:=c.Primes(context.Background(),req)

	if err!=nil{
		log.Fatalf("Error while calling Prines:%v\n",err)
	}

	for{
		res,err:=stream.Recv()

		if err==io.EOF{
			break
		}

		if err!=nil{
			log.Fatalf("Error while reading stream :%v\n",err)
		}
		log.Printf("Primes: %d\n",res.Result)
	}
}