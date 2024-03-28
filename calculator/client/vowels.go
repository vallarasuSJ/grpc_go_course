package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
) 

func doVowelsCount(c pb.CalculatorServiceClient){
	log.Println("doVowelsCount function was invoked") 

	stream,err:=c.VowelsCount(context.Background())

	if err!=nil{
		log.Fatalf("Error while opening stream: %v\n",err)
	} 

	waitc:=make(chan struct{})  

	log.Println("Enter the array of strings") 
	var str[8]string
	go func(){
		for i:=0;i<len(str);i++{
			fmt.Scanln(&str[i])
		}

		for _,s:=range str{
			stream.Send(&pb.VowelRequest{
				Text: s,
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

			log.Println("Vowels count: ",res.Result)
		}
		close(waitc)
	}()

	<-waitc

}