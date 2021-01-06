package main

import (
	"log"
	"net"

	"github.com/daniilsolovey/grpc-service/pkg/adder"
	"github.com/daniilsolovey/grpc-service/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	adderServer := adder.NewAdderServer()
	api.RegisterAdderServer(grpcServer, adderServer)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
