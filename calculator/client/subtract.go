package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
)

func doSubtract(c pb.CalculatorServiceClient){
	log.Println("doSubtract function was invoked") 

	var firstNumber,secondNumber int32
	log.Println("Enter the numbers to subtract:") 
	log.Println("Enter the firstNumber:") 
	_,err:=fmt.Scanln(&firstNumber) 
	if err!=nil{
		log.Fatalf("Error while getting input from the client:%v\n",err)
	}
	log.Println("Enter the secondNumber:") 
	if _,err=fmt.Scanln(&secondNumber); err!=nil{
		log.Fatalf("Error while getting input from the client:%v\n",err)
	}
	req:=&pb.SubRequest{
		FirstNumber: firstNumber,
		SecondNumber: secondNumber,
	}

	res,err:=c.Subtract(context.Background(),req) 
	
	if err!=nil{
		log.Fatalf("Error while receiving data from the server")
	}

	log.Printf("Subtract value is %v\n",res.Result)
}