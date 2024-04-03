package main

import (
	"fmt"
	"log"

	pb "github.com/vallarasuSJ/grpc_go_course/student/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewStudentServiceClient(conn)

	var number int32

	for {
		log.Println("Enter the number to do operation: \n1.Create\n2.Read\n3.Update\n4.Delete\n5.list\n6.Quit")
		fmt.Scanln(&number)
		switch number {
		case 1:
			id := createStudent(c)
			fmt.Println(id)
			fmt.Println("---------------------------------------------------------------------------")
		case 2:
			readStudent(c)
			fmt.Println("---------------------------------------------------------------------------")
		case 3:
			updateStudent(c)
			fmt.Println("---------------------------------------------------------------------------")
		case 4:
			deleteStudent(c)
			fmt.Println("---------------------------------------------------------------------------")
		case 5:
			listStudents(c)
			fmt.Println("---------------------------------------------------------------------------")

		case 6:
			log.Fatal("Exited...")

		default:
			fmt.Println("Enter the correct input to proceed:")
		}
		

	}

	// id:=createStudent(c)

	// readStudent(c,id)

	// updateStudent(c,id)

	// listStudents(c)

	// deleteStudent(c,id)

}
