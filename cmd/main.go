package main

import (
	"log"
	"net"

	"go-grpc/cmd/config"
	"go-grpc/cmd/services"
	produkPb "go-grpc/pb/produk"

	"google.golang.org/grpc"
)

const(
	port = ":50051"
)

func main() {
	// listening port
	netlisten, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to %v", err.Error())
	}

	// Connect Database
	database := config.ConnectDatabase()

	// Set Grpc Server
	grpcServer := grpc.NewServer()
	produkService := services.ProdukService{DB: database}
	produkPb.RegisterProductClientServer(grpcServer, &produkService)

	log.Println("Server started at", netlisten.Addr())
	if err := grpcServer.Serve(netlisten); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}