package app

import (
	"context"

	"cloud.google.com/go/bigtable"
	"github.com/IrisVR/kronos/internal/configuration"
	"github.com/IrisVR/kronos/internal/db"
	"github.com/IrisVR/kronos/internal/grpc"
	log "github.com/sirupsen/logrus"
)

type App struct {
	GrpcServer     grpc.Server
	DatabaseClient *bigtable.Client
}

func NewApp(
	ctx context.Context,
	config configuration.Config,
	databaseClient *bigtable.Client,
	adminClient *bigtable.AdminClient) *App {

	grpcServer := grpc.NewServer(
		config,
		databaseClient,
		adminClient,
	)

	if err := db.EnsureAllTablesExist(ctx, adminClient); err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	return &App{
		DatabaseClient: databaseClient,
		GrpcServer:     grpcServer,
	}
}
