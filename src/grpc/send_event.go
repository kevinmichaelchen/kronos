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

func (s *Server) ReadEvents(ctx context.Context, in *proto.Empty) (*proto.Empty, error) {
	client := s.DatabaseClient
	tbl := client.Open(tableName)

	rowRange := bigtable.PrefixRange("")

	var rowCount int

	var readOpts []bigtable.ReadOption

	// TODO should we use a family filter?
	//readOpts = append(readOpts, bigtable.RowFilter(bigtable.FamilyFilter("links")))

	err := tbl.ReadRows(ctx, rowRange, func(r bigtable.Row) bool {

		rowCount += 1

		log.Infof("Reading columns in row: %s", r.Key())

		// The first one is most recent!
		rowMap := map[string][]bigtable.ReadItem(r)

		for family := range rowMap {
			readItems := rowMap[family]
			// mutations don't actually overwrite. they really append.
			// so the cell is multi-valued until compaction occurs.
			// thus, to get the most recent value, just get the first read item
			//if len(readItems) > 0 {
			//	readItem := readItems[0]
			//	log.Infof("Reading col / time / value: %s / %s / %s",
			//		readItem.Column, readItem.Timestamp.Time(), string(readItem.Value))
			//}
			for _, readItem := range readItems {
				log.Infof("Reading col / time / value: %s / %s / %s",
					readItem.Column, readItem.Timestamp.Time(), string(readItem.Value))
			}
		}

		return true // keep going
	}, readOpts...)
	if err != nil {
		return nil, err
	}

	log.Infof("Found %d rows", rowCount)

	return &proto.Empty{}, nil
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

	return &proto.Empty{}, nil
}
