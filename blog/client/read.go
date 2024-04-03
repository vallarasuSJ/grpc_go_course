package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/blog/proto"
)

func readBlog(c pb.BlogServiceClient,id string) *pb.Blog{ 

	log.Printf("ReadBlog was invoked") 

	req:=&pb.BlogId{Id: id} 

	res,err:=c.ReadBlog(context.Background(),req) 

	if err!=nil{
		log.Printf("Error happened while reading:%v\n",err) 
	}

	log.Printf("Blog was read: %v\n",res)

	return res
}