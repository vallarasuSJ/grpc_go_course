package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/blog/proto"
) 

func UpdateBlog(c pb.BlogServiceClient,id string){
	log.Println("Update blog was invoked") 

	newBlog:=&pb.Blog{
		Id: id,
		AuthorId: "Vallarasu",
		Title: "Hidden mystery", 
		Content: "Thriller content!",
	}

	_,err:=c.UpdateBlog(context.Background(),newBlog)

	if err!=nil{
		log.Fatalf("Error happened while updating : %v\n",err)
	}

	log.Println("Blog was updated")

}