package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListStudents(in *emptypb.Empty,stream pb.StudentService_ListStudentsServer) error {
	log.Println("-----List students was invoked----")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &StudentItem{}

		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB :%v", err),
			)
		} 
		stream.Send(documentToStudent(data))

		if err = cur.Err(); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Unknown internal error:%v", err),
			)
		}
	}
	return nil

}
