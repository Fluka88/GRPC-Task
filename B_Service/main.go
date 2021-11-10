package main

import (
	"B_Service/model"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "myTask/my_task"

	_ "github.com/go-sql-driver/mysql"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDBServiceServer(s, &model.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
