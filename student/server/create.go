package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *Server) CreateStudent(ctx context.Context,in *pb.StudentDetails) (*pb.StudentId, error){
	log.Printf("Create student was invoked %v\n",in) 

	data:=StudentItem{
		StudentName: in.Name,
		Age: in.Age,
	}
	res, err:=collection.InsertOne(ctx,data)

	if err!=nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n",err),
		)
	} 

	oid,ok:=res.InsertedID.(primitive.ObjectID)

	if !ok{
		return nil, status.Errorf(
			codes.Internal,
			"cannot convert to OID",
		)
	}

	return &pb.StudentId{
		Id: oid.Hex(),
	},nil
}