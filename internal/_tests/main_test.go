package _tests

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/IrisVR/kronos/internal/app"
	"github.com/IrisVR/kronos/internal/configuration"
	"github.com/IrisVR/kronos/internal/db"
	log "github.com/sirupsen/logrus"
)

var testConfig configuration.Config
var serverAddress string

func TestMain(m *testing.M) {
	testConfig = configuration.LoadConfig()

	ctx := context.Background()

	client := db.GetBigtableClient(ctx, testConfig)
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("could not close client: %v", err)
		}
	}()

	adminClient := db.GetBigtableAdminClient(ctx, testConfig)
	defer func() {
		if err := adminClient.Close(); err != nil {
			log.Fatalf("could not close client: %v", err)
		}
	}()

	a := app.NewApp(ctx, testConfig, client, adminClient)

	// Create a WaitGroup, which waits for a collection of goroutines to finish
	var wg sync.WaitGroup

	wg.Add(1)
	go a.GrpcServer.Run()

	serverAddress = fmt.Sprintf("localhost:%d", testConfig.GrpcPort)

	// Collect the test code
	code := m.Run()

	// Kill the server
	wg.Done()

	// Exit the program
	os.Exit(code)
}
