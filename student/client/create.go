package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
) 

func createStudent(c pb.StudentServiceClient)string{
	log.Println("Create student was invoked")  
	var studentName,age string
	log.Println("Enter the studentDetails")
	fmt.Println("---------------------------------------------------------------------------")
	log.Println("Enter the Student name: ")
	fmt.Scanln(&studentName)
	log.Println("Enter the Student age: ")
	fmt.Scanln(&age)
	student:=&pb.StudentDetails{
		Name: studentName,
		Age: age,
	}

	res, err:=c.CreateStudent(context.Background(),student) 

	if err!=nil{
		log.Fatalf("Unexpected error: %v\n",err)
	}
	log.Printf("Student has been created: %s\n",res.Id)
	return res.Id
}