package main

import (
	"log"
	"net"

	pb "github.com/vallarasuSJ/grpc_go_course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" 
	"google.golang.org/grpc/reflection"
)
 
var addr string = "0.0.0.0:50051" 

type Server struct{
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp",addr) 

	if err!=nil{
		log.Fatalf("Failed to listen on %v\n",err)
	}
	log.Printf("Listening on %s\n",addr) 

	opts:=[]grpc.ServerOption{}
	tls:=true  //change that to false if needed 

	if tls{
		certFile:="ssl/server.crt"
		keyFile:="ssl/server.pem"
		creds,err:=credentials.NewServerTLSFromFile(certFile,keyFile) 

		if err!=nil{
			log.Fatalf("Failed to loading certificates: %v\n",err)
		}  

		opts=append(opts, grpc.Creds(creds))


	}

	s:=grpc.NewServer(opts...) 
	pb.RegisterCalculatorServiceServer(s,&Server{}) 
	reflection.Register(s)

	if err = s.Serve(lis); err!=nil{
		log.Fatalf("Failed to serve: %v\n",err)
	}

}