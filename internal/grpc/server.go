package grpc

import (
	"fmt"

	"cloud.google.com/go/bigtable"
	"github.com/IrisVR/kronos/internal/configuration"
	proto "github.com/IrisVR/kronos/internal/pb"
	"github.com/IrisVR/nucleus/kontext"
	"google.golang.org/grpc/codes"

	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
)

// Server represents a gRPC server. It handles all gRPC calls.
type Server struct {
	Port           int
	DatabaseClient *bigtable.Client
	AdminClient    *bigtable.AdminClient
}

func NewServer(
	config configuration.Config,
	databaseClient *bigtable.Client,
	adminClient *bigtable.AdminClient,
) Server {
	return Server{
		Port:           config.GrpcPort,
		DatabaseClient: databaseClient,
		AdminClient:    adminClient,
	}
}

func customCodeToLogLevel(code codes.Code) logrus.Level {
	switch code {
	case codes.OK:
		return logrus.TraceLevel
	case codes.Canceled:
		return logrus.InfoLevel
	case codes.Unknown:
		return logrus.ErrorLevel
	case codes.InvalidArgument:
		return logrus.InfoLevel
	case codes.DeadlineExceeded:
		return logrus.WarnLevel
	case codes.NotFound:
		return logrus.InfoLevel
	case codes.AlreadyExists:
		return logrus.InfoLevel
	case codes.PermissionDenied:
		return logrus.WarnLevel
	case codes.Unauthenticated:
		return logrus.InfoLevel // unauthenticated requests can happen
	case codes.ResourceExhausted:
		return logrus.WarnLevel
	case codes.FailedPrecondition:
		return logrus.WarnLevel
	case codes.Aborted:
		return logrus.WarnLevel
	case codes.OutOfRange:
		return logrus.WarnLevel
	case codes.Unimplemented:
		return logrus.ErrorLevel
	case codes.Internal:
		return logrus.ErrorLevel
	case codes.Unavailable:
		return logrus.WarnLevel
	case codes.DataLoss:
		return logrus.ErrorLevel
	default:
		return logrus.ErrorLevel
	}
}

// Run starts the gRPC server.
func (s *Server) Run() {
	address := fmt.Sprintf(":%d", s.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}
	logrus.Printf("Starting gRPC server on %s...\n", address)

	server := kontext.NewServer()

	// Register our services
	proto.RegisterEventServiceServer(server, s)

	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
