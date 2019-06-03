package app

import (
	"cloud.google.com/go/bigtable"
	"context"
	"github.com/IrisVR/kronos/configuration"
	"github.com/IrisVR/kronos/db"
	"github.com/IrisVR/kronos/grpc"
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
