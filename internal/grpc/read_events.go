package grpc

import (
	"context"

	"cloud.google.com/go/bigtable"
	"github.com/IrisVR/kronos/internal/db"
	proto "github.com/IrisVR/kronos/internal/pb"
	log "github.com/sirupsen/logrus"
)

func (s *Server) ReadEvents(ctx context.Context, in *proto.Empty) (*proto.Empty, error) {
	if err := s.readAllRows(ctx, db.LoginTable); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *Server) readAllRows(ctx context.Context, tblName string) error {
	client := s.DatabaseClient
	tbl := client.Open(tblName)
	rowRange := bigtable.PrefixRange("")

	var rowCount int

	var readOpts []bigtable.ReadOption

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
		return err
	}

	log.Infof("Found %d rows in table: %s", rowCount, tblName)
	return nil
}
