package db

import (
	"cloud.google.com/go/bigtable"
	"context"
	"github.com/IrisVR/kronos/configuration"
	log "github.com/sirupsen/logrus"
)

const (
	LoginTable       = "logins"
	LoginFamily      = "logins"
	HeartbeatsTable  = "heartbeats"
	HeartbeatsFamily = "heartbeats"
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
	if columnFamily == "" {
		return nil
	}
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

func EnsureAllTablesExist(ctx context.Context, client *bigtable.AdminClient) error {
	if err := ensureTableExists(ctx, client, LoginTable); err != nil {
		return err
	}
	if err := ensureTableExists(ctx, client, HeartbeatsTable); err != nil {
		return err
	}
	return nil
}

func ensureTableExists(ctx context.Context, client *bigtable.AdminClient, tableName string) error {
	log.Infof("Ensuring table '%s' exists", tableName)
	if _, err := client.TableInfo(ctx, tableName); err != nil {
		log.Infof("Table '%s' does not exist. Creating it...", tableName)
		if err := client.CreateTable(ctx, tableName); err != nil {
			return err
		}
		log.Infof("Created table: %s", tableName)
	} else {
		log.Infof("Table exists: '%s'", tableName)
	}
	return nil
}
