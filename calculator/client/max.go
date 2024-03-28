package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient){
	log.Println("doMax was invoked")  

	stream,err:=c.Max(context.Background())

	if err!=nil{
		log.Fatalf("Error while opening stream: %v\n",err)
	} 

	waitc:=make(chan struct{})

	go func(){
		numbers:=[]int32{3,4,3,32,3,24,5,6,4,87,98} 

		for _,number:=range numbers{
			log.Printf("Sending number: %d\n",number) 
			stream.Send(&pb.MaxRequest{
				Number:number,
			})
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()

	go func(){
		for{
			res,err:=stream.Recv() 

			if err==io.EOF{
				break
			}

			if err!=nil{
				log.Fatalf("Problem while reading server stream: %v",err)
			}
			log.Printf("Received a new maximum: %d\n",res.Result)
		}
		close(waitc)
	}()

	<-waitc


}