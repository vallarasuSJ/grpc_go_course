package main

import (
	"context"
	"fmt"

	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
) 

func doMaxEvenNumber(c pb.CalculatorServiceClient){ 
	log.Println("doMaxEvenNumber was invoked")

	arr:=[]int32{53,3,43,5,34,5,6,4,57,54,6}
	stream,err:=c.MaxEvenNumber(context.Background())
	if err!=nil{
		log.Fatalf("Error while opening the stream %v\n",err)
	}

	log.Println("Enter the array of numbers: ")
	
	for i:=int32(0);i<10;i++{
		fmt.Scanln(&arr[i])
	}
	
	for _,num:=range arr{
		stream.Send(&pb.EvenRequest{
			EvenNumber: num,
		})
	}

	res,err:=stream.CloseAndRecv() 

	if err!=nil{
		log.Fatalf("Error while receiving response:%v\n",err) 
	}

	log.Printf("Maximum even number: %d\n",res.Result)

}