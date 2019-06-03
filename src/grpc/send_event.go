package grpc

import (
	"cloud.google.com/go/bigtable"
	"context"
	"fmt"
	"github.com/IrisVR/kronos/db"
	proto "github.com/IrisVR/kronos/pb"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	loginTable      = "logins"
	heartbeatsTable = "heartbeats"
)

func (s *Server) SendHeartbeatEvent(ctx context.Context, in *proto.Event) (*proto.Empty, error) {
	// TODO validate event, e.g.,v4 uuid
	if err := s.writeEvent(ctx, heartbeatsTable, in); err != nil {
		return nil, err
	}

	log.Infof("Processed a heartbeat event")

	return &proto.Empty{}, nil
}

func (s *Server) SendLoginEvent(ctx context.Context, in *proto.Event) (*proto.Empty, error) {
	// TODO validate event, e.g.,v4 uuid
	if err := s.writeEvent(ctx, loginTable, in); err != nil {
		return nil, err
	}

	log.Infof("Processed a login event")

	return &proto.Empty{}, nil
}

func (s *Server) writeEvent(ctx context.Context, tableName string, in *proto.Event) error {
	client := s.DatabaseClient
	tbl := client.Open(tableName)

	eventTime := getTime(in.TimeMs)
	rowKey := getLoginRowKey(in.UserID, in.TimeMs)
	columnFamily := in.UserID
	columnName := "value"

	// Create the column family
	if err := db.EnsureFamilyExists(ctx, s.AdminClient, tableName, columnFamily); err != nil {
		return err
	}

	mut := bigtable.NewMutation()
	mut.Set(columnFamily, columnName, bigtable.Time(eventTime), []byte("1"))
	return tbl.Apply(ctx, rowKey, mut)
}

func getTime(epochMs int64) time.Time {
	s := epochMs / 1000
	ns := (epochMs % 1000) * 1000000
	return time.Unix(s, ns)
}

func reverseTimestamp(epochMs int64) string {
	// TODO actually reverse it.
	//  for now we don't care about hotspotting
	return strconv.Itoa(int(epochMs))
}

func getLoginRowKey(userID string, epochMS int64) string {
	return fmt.Sprintf("%s:%s", userID, reverseTimestamp(epochMS))
}
