package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
) 

func updateStudent(c pb.StudentServiceClient){ 

	log.Println("update Student was invoked")

	var id,name,age string 
	log.Println("Enter the student details to be updated")
	fmt.Println("---------------------------------------------------------------------------")
	log.Println("Enter the id:")
	fmt.Scanln(&id)
	log.Println("Enter the student name:")
	fmt.Scanln(&name)
	log.Println("Enter the id:")
	fmt.Scanln(&age)

	newStudent:=&pb.StudentDetails{
		Id: id,
		Name: name,
		Age: age,
	}

	_,err:=c.UpdateStudent(context.Background(),newStudent) 

	if err!=nil{
		log.Fatalf("Error happened while updating : %v\n",err)
	} 

	log.Println("Student was updated")

}