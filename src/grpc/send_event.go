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

// TODO validate event, e.g.,v4 uuid

func (s *Server) SendHeartbeatEvent(ctx context.Context, in *proto.Event) (*proto.Empty, error) {
	rowKey := getHeartbeatRowKey(in.UserID, in.TimeMs)
	// Column family cannot be empty string
	columnFamily := db.HeartbeatsFamily
	eventTime := bigtable.Time(getTime(in.TimeMs))

	if err := s.writeEvent(ctx, db.HeartbeatsTable, rowKey, columnFamily, eventTime); err != nil {
		return nil, err
	}

	log.Infof("Processed a heartbeat event")

	return &proto.Empty{}, nil
}

func (s *Server) SendLoginEvent(ctx context.Context, in *proto.Event) (*proto.Empty, error) {
	rowKey := getLoginRowKey(in.UserID, in.TimeMs)
	// Column family cannot be empty string
	columnFamily := db.LoginFamily
	eventTime := bigtable.Time(getTime(in.TimeMs))

	if err := s.writeEvent(ctx, db.LoginTable, rowKey, columnFamily, eventTime); err != nil {
		return nil, err
	}

	log.Infof("Processed a login event")

	return &proto.Empty{}, nil
}

func (s *Server) writeEvent(
	ctx context.Context,
	tableName, rowKey, columnFamily string,
	t bigtable.Timestamp,
) error {
	client := s.DatabaseClient
	columnName := "value"

	tbl := client.Open(tableName)

	if err := db.EnsureFamilyExists(ctx, s.AdminClient, tableName, columnFamily); err != nil {
		return err
	}

	mut := bigtable.NewMutation()
	mut.Set(columnFamily, columnName, t, []byte("1"))
	return tbl.Apply(ctx, rowKey, mut)
}

func getTime(epochMs int64) time.Time {
	s := epochMs / 1000
	ns := (epochMs % 1000) * 1000000
	return time.Unix(s, ns)
}

func stringifyTimestamp(epochMs int64) string {
	return strconv.Itoa(int(epochMs))
}

func reverseTimestamp(epochMs int64) string {
	return reverseString(stringifyTimestamp(epochMs))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getLoginRowKey(userID string, epochMS int64) string {
	return fmt.Sprintf("%s:%s", userID, reverseTimestamp(epochMS))
}

func getHeartbeatRowKey(userID string, epochMS int64) string {
	return fmt.Sprintf("%s:%s", userID, stringifyTimestamp(epochMS))
}
