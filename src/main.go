package main

import (
	"cloud.google.com/go/bigtable"
	"context"
	"github.com/IrisVR/kronos/app"
	"github.com/IrisVR/kronos/configuration"
	log "github.com/sirupsen/logrus"
	"sync"
)

func main() {
	config := configuration.LoadConfig()

	ctx := context.Background()

	client := getBigtableClient(ctx, config)
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("could not close client: %v", err)
		}
	}()

	a := app.NewApp(config, client)

	// Create a WaitGroup, which waits for a collection of goroutines to finish
	var wg sync.WaitGroup

	// Run the gRPC server
	wg.Add(1)
	go a.GrpcServer.Run()

	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
}

func getBigtableClient(ctx context.Context, config configuration.Config) *bigtable.Client {
	project := "my-project"
	instance := "my-instance"
	client, err := bigtable.NewClient(ctx, project, instance)
	if err != nil {
		log.Fatalf("could not create Bigtable client: %v", err)
	}
	return client
}
