package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/vallarasuSJ/grpc_go_course/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect:%v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)

	readBlog(c, id)
	//    readBlog(c,"NonExisting id")
	UpdateBlog(c, id)

	listBlog(c)
  
	deleteBlog(c,id)
	

	//mongo localhost:8081 - user:admin , pass:pass
}
