package app

import (
	"cloud.google.com/go/bigtable"
	"github.com/IrisVR/kronos/configuration"
	"github.com/IrisVR/kronos/grpc"
)

type App struct {
	GrpcServer     grpc.Server
	DatabaseClient *bigtable.Client
}

func NewApp(
	config configuration.Config,
	databaseClient *bigtable.Client,
	adminClient *bigtable.AdminClient) *App {

	grpcServer := grpc.NewServer(
		config,
		databaseClient,
		adminClient,
	)

	return &App{
		DatabaseClient: databaseClient,
		GrpcServer:     grpcServer,
	}
}
