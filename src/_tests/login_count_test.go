package _tests

import (
	"context"
	proto "github.com/IrisVR/kronos/pb"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"testing"
)

func Test_GetLoginCount(t *testing.T) {
	///////////////////////////////////
	// Create a client to perform reads
	///////////////////////////////////
	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Could not create gRPC connection: %v", err)
	}
	client := proto.NewEventServiceClient(connection)

	userID := uuid.Must(uuid.NewRandom()).String()

	Convey("When we create login events", t, func(c C) {
		ctx := context.TODO()

		res, err := client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 0)

		_, err = client.SendLoginEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 1559520445749,
		})
		So(err, ShouldBeNil)

		res, err = client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 1)

		res, err = client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
			Start:  1559520445740,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 1)

		_, err = client.SendLoginEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 1559520445752,
		})
		So(err, ShouldBeNil)

		_, err = client.SendLoginEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 1559520445753,
		})
		So(err, ShouldBeNil)

		res, err = client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
			Start:  1559520445740,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 3)
	})
}
