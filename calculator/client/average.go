package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
) 

func doAverage(c pb.CalculatorServiceClient){
	log.Println("doAvg was invoked")

	stream,err:=c.Average(context.Background())

	if err!=nil{
		log.Fatalf("Error while opening the stream %v\n",err)
	}
	numbers:=[]int32{3,4,5,3,43}

	for _,number:=range numbers{
		log.Printf("Sending number:%d \n",number)
		stream.Send(&pb.AverageRequest{
			Number: int64(number),
		})
	}

	res,err:=stream.CloseAndRecv() 

	if err!=nil{
		log.Fatalf("Error while receiving response:%v\n",err) 
	}

	log.Printf("Average: %f\n",res.Result)

}