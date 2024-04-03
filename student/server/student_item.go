package main

import (
	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentItem struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`  
	StudentName string `bson:"student_name"` 
	Age string `bson:"student_age"`
}  

func documentToStudent(data *StudentItem) *pb.StudentDetails{
	return &pb.StudentDetails{
		Id: data.ID.Hex(),
		Name: data.StudentName,
		Age: data.Age,
	}
}

