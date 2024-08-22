package main

import (
	"fmt"
	"grpc-users/configs"
	pb "grpc-users/pb"
	"grpc-users/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	pgConnection, err := configs.PosgreConnection()
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
	}
	db, err := pgConnection.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}
	defer db.Close()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9002))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &services.UserService{Db: pgConnection})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
