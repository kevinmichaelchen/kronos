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

// EnsureFamilyExists creates the column family if one doesn't already exist.
func EnsureFamilyExists(ctx context.Context, client *bigtable.AdminClient, tableName, columnFamily string) error {
	var familyExists bool
	if tableInfo, err := client.TableInfo(ctx, tableName); err != nil {
		return err
	} else {
		familyInfos := tableInfo.FamilyInfos
		for _, familyInfo := range familyInfos {
			if familyInfo.Name == columnFamily {
				familyExists = true
				break
			}
		}
	}
	if !familyExists {
		if err := client.CreateColumnFamily(ctx, tableName, columnFamily); err != nil {
			return err
		}
	}
	return nil
}
