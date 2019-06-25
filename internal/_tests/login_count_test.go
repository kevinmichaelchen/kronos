package _tests

import (
	"context"
	"testing"

	proto "github.com/IrisVR/kronos/internal/pb"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
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

		// There should not be any logins for this user yet
		res, err := client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 0)

		// Send a login event
		_, err = client.SendLoginEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 1559520445749,
		})
		So(err, ShouldBeNil)

		// Okay, there should now be 1 login event for this user since the dawn of time
		res, err = client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 1)

		// There should be one login since a certain start time
		res, err = client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
			Start:  1559520445740,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 1)

		// Send another login event
		_, err = client.SendLoginEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 1559520445752,
		})
		So(err, ShouldBeNil)

		// Send another login event
		_, err = client.SendLoginEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 1559520445753,
		})
		So(err, ShouldBeNil)

		// There should now be 3 login events since a certain start time
		res, err = client.GetNumberOfLogins(ctx, &proto.UserQuery{
			UserID: userID,
			Start:  1559520445740,
		})
		So(err, ShouldBeNil)
		So(res.Count, ShouldEqual, 3)
	})
}
