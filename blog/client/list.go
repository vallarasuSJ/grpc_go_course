package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
) 

func listBlog(c pb.BlogServiceClient){
	log.Println("List blog was invoked") 

	stream,err:=c.ListBlogs(context.Background(),&emptypb.Empty{})

	if err!=nil{
		log.Fatalf("Error while calling ListBlogs: %v\n",err)
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