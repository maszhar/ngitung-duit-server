package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/djeniusinvfest/inventora/auth/handler"
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	errorLog := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		errorLog.Printf("No .env file. Using system environment")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("[::]:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthenticationServer(server, handler.NewHandler())

	log.Printf("server listening at: %s", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
