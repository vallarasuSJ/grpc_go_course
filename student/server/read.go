package main

import (
	"context"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func(s *Server) ReadStudent(ctx context.Context, in *pb.StudentId) (*pb.StudentDetails, error){
	log.Printf("Read student invoked with: %v\n",in) 
	
	oid,err:=primitive.ObjectIDFromHex(in.Id) 

	if err!=nil{
		return nil,status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data:=&StudentItem{}
	filter:=bson.M{"_id":oid}

	res:=collection.FindOne(ctx,filter) 

	if err:=res.Decode(data);err!=nil{
		return nil,status.Errorf(
			codes.NotFound,
			"Cannot find student with the id provided",
		)
	} 

	return documentToStudent(data),nil

	
}