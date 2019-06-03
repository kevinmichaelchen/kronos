package grpc

import (
	"cloud.google.com/go/bigtable"
	"context"
	"github.com/IrisVR/kronos/db"
	proto "github.com/IrisVR/kronos/pb"
	log "github.com/sirupsen/logrus"
)

func (s *Server) GetNumberOfLogins(ctx context.Context, in *proto.UserQuery) (*proto.CountResponse, error) {
	client := s.DatabaseClient
	tbl := client.Open(db.LoginTable)
	rowRange := bigtable.PrefixRange(in.UserID)

	var rowCount int

	var readOpts []bigtable.ReadOption

	// Filter on the column family
	//readOpts = append(readOpts, bigtable.RowFilter(bigtable.FamilyFilter(db.LoginFamily)))

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
	return &proto.DurationResponse{
		DurationMs: 0,
	}, nil
}
