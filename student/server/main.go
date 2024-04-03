
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)
var collection *mongo.Collection

var addr string = "0.0.0.0:50051" 

type Server struct{
	pb.StudentServiceServer
}

func main() { 
	client, err:=mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/")) 

	if err!=nil{
		log.Fatal(err)
	}  

	err=client.Connect(context.Background())

	if err!=nil{
		log.Fatal(err)
	}
	collection=client.Database("studentdb").Collection("student")

	lis, err := net.Listen("tcp",addr) 

	if err!=nil{
		log.Fatalf("Failed to listen on: %v\n",err)
	}

	log.Printf("Listening on %s\n",addr)   
	
	s:=grpc.NewServer()
	pb.RegisterStudentServiceServer(s,&Server{})

	if err=s.Serve(lis); err!=nil{
		log.Fatalf("Failed to serve : %v\n",err)
	}

}