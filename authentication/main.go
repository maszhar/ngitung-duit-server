package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/djeniusinvfest/inventora/auth/handler"
	"github.com/djeniusinvfest/inventora/auth/model"
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	errorLog := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		errorLog.Printf("No .env file. Using system environment")
	}

	// Setup listening port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Connect to database
	mongoUri := os.Getenv("MONGO_URI")
	dbClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	db := dbClient.Database("auth")

	// Init models
	userModel := model.NewUserModel(db)

	// Init repositories
	authRepo := repository.NewAuthRepo(userModel)

	listener, err := net.Listen("tcp", fmt.Sprintf("[::]:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthenticationServer(server, handler.NewHandler(authRepo))

	log.Printf("server listening at: %s", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
