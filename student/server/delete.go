package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)  

func (s *Server)DeleteStudent(ctx context.Context, in *pb.StudentId) (*emptypb.Empty, error){
	log.Printf("deleteStudent was invoked with: %v\n",in)  

	oid,err:=primitive.ObjectIDFromHex(in.Id) 

	if err!=nil{
		return nil,status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	res,err:=collection.DeleteOne(ctx,bson.M{"_id":oid}) 

	if err!=nil{
		return nil,status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v\n",err),
		)
	}
	
	if res.DeletedCount==0{
		return nil,status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	} 

	return &emptypb.Empty{},nil
	

}