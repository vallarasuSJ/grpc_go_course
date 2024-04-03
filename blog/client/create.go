package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/blog/proto"
) 

func createBlog(c pb.BlogServiceClient)string{
	log.Println("Create blog was invoked") 

	blog:=&pb.Blog{
		AuthorId: "Clement",
		Title: "My first blog", 
		Content: "Content of the first blog",
	}

	res, err:=c.CreateBlog(context.Background(),blog) 

	if err!=nil{
		log.Fatalf("Unexpected error: %v\n",err)
	}
	log.Printf("Blog has been created: %s\n",res.Id)
	return res.Id
}