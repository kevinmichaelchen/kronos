package _tests

import (
	"context"
	proto "github.com/IrisVR/kronos/pb"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"testing"
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

	userID := uuid.Must(uuid.NewRandom()).String()

	Convey("When we send heartbeats", t, func(c C) {
		ctx := context.TODO()

		// There should not be any session activity for this user yet
		res, err := client.GetUserSessionDuration(ctx, &proto.UserQuery{
			UserID: userID,
		})
		So(err, ShouldBeNil)
		So(res.DurationMs, ShouldEqual, 0)

		// TODO in theory every heartbeat should equal 5 seconds
		//  but some heartbeats might fail to send...
		//  so we should only increment the duration if the
		//  current heartbeat's time is 5 seconds greater than the previous one's.
		//  Yes, this means when iterating we have to store the previous heartbeat.
	})
}
