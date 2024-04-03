package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"google.golang.org/protobuf/types/known/emptypb"
) 

func listStudents(c pb.StudentServiceClient){
	log.Println("----listStudents was invoked----")  

	stream,err:=c.ListStudents(context.Background(),&emptypb.Empty{})

	if err!=nil{
		log.Fatalf("Error while calling list Students: %v\n",err)
	}

	for{
		res,err:=stream.Recv() 

		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("Something happened: %v\n",err)
		}
		log.Println(res)
	}
	
}