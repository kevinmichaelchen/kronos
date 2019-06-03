package db

import (
	"cloud.google.com/go/bigtable"
	"context"
	"github.com/IrisVR/kronos/configuration"
	log "github.com/sirupsen/logrus"
)

func GetBigtableClient(ctx context.Context, config configuration.Config) *bigtable.Client {
	project := config.BigtableProject
	instance := config.BigtableInstance
	client, err := bigtable.NewClient(ctx, project, instance)
	if err != nil {
		log.Fatalf("could not create Bigtable client: %v", err)
	}
	return client
}

func GetBigtableAdminClient(ctx context.Context, config configuration.Config) *bigtable.AdminClient {
	project := config.BigtableProject
	instance := config.BigtableInstance
	client, err := bigtable.NewAdminClient(ctx, project, instance)
	if err != nil {
		log.Fatalf("could not create Bigtable admin client: %v", err)
	}
	return client
}
