package grpc

import (
	"context"

	"cloud.google.com/go/bigtable"
	"github.com/IrisVR/kronos/internal/db"
	proto "github.com/IrisVR/kronos/internal/pb"
	log "github.com/sirupsen/logrus"
)

func (s *Server) GetNumberOfLogins(ctx context.Context, in *proto.UserQuery) (*proto.CountResponse, error) {
	client := s.DatabaseClient
	tbl := client.Open(db.LoginTable)
	rowRange := bigtable.PrefixRange(in.UserID)

	var rowCount int
	var readOpts []bigtable.ReadOption

	err := tbl.ReadRows(ctx, rowRange, func(r bigtable.Row) bool {

		log.Infof("Reading columns in row: %s", r.Key())
		rowCount += 1

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

	return &proto.CountResponse{
		Count: int64(rowCount),
	}, nil
}

func (s *Server) GetUserSessionDuration(ctx context.Context, in *proto.UserQuery) (*proto.DurationResponse, error) {
	// TODO in theory every heartbeat should equal 5 seconds
	//  but some heartbeats might fail to send...
	//  so we should only increment the duration if the
	//  current heartbeat's time is 5 seconds greater than the previous one's.
	//  Yes, this means when iterating we have to store the previous heartbeat.

	client := s.DatabaseClient
	tbl := client.Open(db.HeartbeatsTable)

	prefix := in.UserID

	rowRange := bigtable.PrefixRange(prefix)

	var rowCount int
	var readOpts []bigtable.ReadOption

	readOpts = append(readOpts, bigtable.RowFilter(bigtable.FamilyFilter(db.HeartbeatsFamily)))

	var internalError error

	err := tbl.ReadRows(ctx, rowRange, func(r bigtable.Row) bool {

		rowKey, err := getHeartbeatRowKeyFromString(r.Key())
		if err != nil {
			internalError = err
			return false
		}

		if in.Start != 0 {
			if rowKey.IsBefore(in.Start) {
				return true
			}
		}

		if in.End != 0 {
			if rowKey.IsAfter(in.End) {
				return false
			}
		}

		log.Infof("Reading columns in row: %s", r.Key())
		rowCount += 1

		// The first one is most recent!
		rowMap := map[string][]bigtable.ReadItem(r)

		for family := range rowMap {
			readItems := rowMap[family]
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

	if internalError != nil {
		return nil, internalError
	}

	return &proto.DurationResponse{
		DurationMs: int64(rowCount * 5),
	}, nil
}
