package main

import (
	"io"
	"log"
	"unicode"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func (s *Server) VowelsCount(stream pb.CalculatorService_VowelsCountServer) error{
	log.Println("VowelsCount function was invoked from the server")

	
	
	for{
		var count int32=0
		str,err:=stream.Recv()  

		if err==io.EOF{
			return nil
		}

		if err!=nil{
			log.Fatalf("Error while reading client stream %v\n",err) 
		}
		ch:=[]rune(str.Text)
		for i:=0;i<len(ch);i++{
			c:=unicode.ToLower(ch[i])
			if c=='a' || c=='e' || c=='i' || c=='o' || c=='u'{
				count++
			}
		}

		err=stream.Send(&pb.VowelResponse{
			Result: count,
		}) 

		if err!=nil{
			log.Fatalf("Error while sending data to client: %v\n",err)
		}
	}
}