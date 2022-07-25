package main

import (
	"log"

	"github.com/dosanma1/go-grpc-wishlist/config"
	"github.com/dosanma1/go-grpc-wishlist/internal/server/grpc"
)

func main() {
	// var logger = logrus.New()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewGrpcServer(cfg)

	log.Fatalln(grpcServer.Run())
}
