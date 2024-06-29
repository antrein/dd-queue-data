package main

import (
	"antrein/dd-queue-data/application/common/repository"
	"antrein/dd-queue-data/application/common/resource"
	"antrein/dd-queue-data/application/grpc"
	"antrein/dd-queue-data/application/rest"
	"antrein/dd-queue-data/model/config"
	"context"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	rsc, err := resource.NewCommonResource(cfg, ctx)
	if err != nil {
		log.Fatal(err)
	}
	repository, err := repository.NewCommonRepository(cfg, rsc)
	if err != nil {
		log.Fatal(err)
	}
	rest_app, err := rest.ApplicationDelegate(cfg, repository)
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server concurrently
	go func() {
		grpc_app, err := grpc.ApplicationDelegate(cfg, repository)
		if err != nil {
			log.Fatal(err)
		}
		if err := grpc.StartServer(cfg, grpc_app); err != nil {
			log.Fatal(err)
		}
	}()

	if err = rest.StartServer(cfg, rest_app); err != nil {
		log.Fatal(err)
	}

}
