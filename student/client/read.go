package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
) 


func readStudent(c pb.StudentServiceClient) *pb.StudentDetails{
	log.Printf("readStudent was invoked")  

	var id string 
	log.Println("Enter the id: ")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Scanln(&id)
	req:=&pb.StudentId{Id: id}

	res,err:=c.ReadStudent(context.Background(),req) 

	if err!=nil{
		log.Printf("Error happened while reading:%v\n",err)
	} 

	log.Printf("Student was read: %v\n",res)
	return res
}