package _tests

import (
	"context"
	"testing"

	proto "github.com/IrisVR/kronos/internal/pb"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
)

func Test_GetSessionDuration(t *testing.T) {
	///////////////////////////////////
	// Create a client to perform reads
	///////////////////////////////////
	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Could not create gRPC connection: %v", err)
	}
	client := proto.NewEventServiceClient(connection)

	userID := uuid.New().String()

	heartbeatPeriodicity := 5

	Convey("When we send heartbeats", t, func(c C) {
		ctx := context.TODO()

		// There should not be any session activity for this user yet
		res, err := client.GetUserSessionDuration(ctx, &proto.UserQuery{
			UserID: userID,
		})
		So(err, ShouldBeNil)
		So(res.DurationMs, ShouldEqual, 0)

		// Send heartbeat
		_, err = client.SendHeartbeatEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 10000,
			Properties: map[string]string{
				"initial": "true",
			},
		})
		So(err, ShouldBeNil)

		// Session activity should still be 0
		res, err = client.GetUserSessionDuration(ctx, &proto.UserQuery{
			UserID: userID,
			Start:  15000,
		})
		So(err, ShouldBeNil)
		So(res.DurationMs, ShouldEqual, 0)

		// Send heartbeat
		_, err = client.SendHeartbeatEvent(ctx, &proto.Event{
			UserID: userID,
			TimeMs: 20000,
		})
		So(err, ShouldBeNil)

		res, err = client.GetUserSessionDuration(ctx, &proto.UserQuery{
			UserID: userID,
			Start:  15000,
		})
		So(err, ShouldBeNil)
		So(res.DurationMs, ShouldEqual, 1*heartbeatPeriodicity)
	})
}
