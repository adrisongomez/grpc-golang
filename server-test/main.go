package main

import (
	"log"
	"net"

	"github.com/adrisongomez/grpc_golang/database"
	"github.com/adrisongomez/grpc_golang/server"
	"github.com/adrisongomez/grpc_golang/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
    list, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	server := server.NewTestServer(repo)

	s := grpc.NewServer()

	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err != nil)
	}

}
