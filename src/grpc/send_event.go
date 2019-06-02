package grpc

import (
	"cloud.google.com/go/bigtable"
	"context"
	"fmt"
	proto "github.com/IrisVR/kronos/pb"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	tableName = "mytable"
)

func getTime(in *proto.Event) time.Time {
	epochMs := in.Time
	s := epochMs / 1000
	ns := (epochMs % 1000) * 1000000
	return time.Unix(s, ns)
}

func getRowKeyFromEvent(ctx context.Context, in *proto.Event, now time.Time) string {
	return fmt.Sprintf("%s#%d", in.Event, now.UnixNano()/1000000)
}

func (s *Server) SendEvent(ctx context.Context, in *proto.Event) (*proto.Empty, error) {
	client := s.DatabaseClient
	tbl := client.Open(tableName)

	// TODO normalize events? check that they're part of a set of known event types

	now := time.Now()
	// TODO an event might need several row keys since we save the same event in a few different ways
	//  to satisfy various queries.
	rowKey := getRowKeyFromEvent(ctx, in, now)
	columnFamily := "event"
	columnName := "value"

	mut := bigtable.NewMutation()
	mut.Set(columnFamily, columnName, bigtable.Time(now), []byte("1"))
	err := tbl.Apply(ctx, rowKey, mut)
	if err != nil {
		return nil, err
	}

	log.Infof("Processed a %s event", in.Event)

	return &proto.Empty{}, nil
}
