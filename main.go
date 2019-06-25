package main

import (
	"context"
	"sync"

	"github.com/IrisVR/kronos/internal/app"
	"github.com/IrisVR/kronos/internal/configuration"
	"github.com/IrisVR/kronos/internal/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := configuration.LoadConfig()

	ctx := context.Background()

	client := db.GetBigtableClient(ctx, config)
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("could not close client: %v", err)
		}
	}()

	adminClient := db.GetBigtableAdminClient(ctx, config)
	defer func() {
		if err := adminClient.Close(); err != nil {
			log.Fatalf("could not close client: %v", err)
		}
	}()

	a := app.NewApp(ctx, config, client, adminClient)

	// Create a WaitGroup, which waits for a collection of goroutines to finish
	var wg sync.WaitGroup

	// Run the gRPC server
	wg.Add(1)
	go a.GrpcServer.Run()

	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
}
