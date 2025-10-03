package main

import (
	"bookstore_grpc/controllers"
	"bookstore_grpc/db"
	"bookstore_grpc/middleware"
	"bookstore_grpc/proto"
	"bookstore_grpc/utils"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	utils.LoadEnv()
	db.InitDB()

	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
	grpc.UnaryInterceptor(middleware.AuthInterceptor),
)
	proto.RegisterBookstoreServiceServer(grpcServer, controllers.NewBookstoreController())

	fmt.Println("gRPC server running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
