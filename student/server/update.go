package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateStudent(ctx context.Context,in *pb.StudentDetails) (*emptypb.Empty, error){
	log.Printf("update student was invoked with %v\n",in)  

	oid,err:=primitive.ObjectIDFromHex(in.Id) 

	if err!=nil{
		return nil,status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data:=&StudentItem{
		StudentName: in.Name,
		Age: in.Age,
	}

	res,err:=collection.UpdateOne(
		ctx,
		bson.M{"_id":oid},
		bson.M{"$set":data},
	)

	if err!=nil{
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}  

	if res.MatchedCount==0{
		return nil,status.Errorf(
			codes.Internal,
			"Cannot find student with id",
		)
	}
	return &emptypb.Empty{},nil

	
}