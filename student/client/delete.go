package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
)

func deleteStudent(c pb.StudentServiceClient){

	log.Printf("deleteStudent was invoked\n") 

	var id string
	log.Printf("Enter the id to be deleted:\n")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Scanln(&id)
	_,err:=c.DeleteStudent(context.Background(),&pb.StudentId{Id: id}) 

	if err!=nil{
		log.Fatalf("Error while deleting: %v\n",err)
	}
	log.Println("student was deleted!!")
}
