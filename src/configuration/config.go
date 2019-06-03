package configuration

import (
	"encoding/json"
	"github.com/IrisVR/nucleus/logging"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagForGrpcPort = "grpc_port"
)

type Config struct {
	// AppID is a unique identifier for the instance (pod) running this app.
	AppID string

	// GrpcPort controls what port our gRPC server runs on.
	GrpcPort int

	// LoggingConfig configures how we do logging.
	LoggingConfig logging.Config

	BigtableProject  string
	BigtableInstance string
}

func (c Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Could not marshal config to string: %v", err)
	}
	return string(b)
}

func LoadConfig() Config {
	c := Config{
		AppID:            uuid.Must(uuid.NewRandom()).String(),
		GrpcPort:         8080,
		BigtableProject:  "my-project",
		BigtableInstance: "my-instance",
	}

	flag.Int(flagForGrpcPort, c.GrpcPort, "port to serve gRPC on")

	flag.Parse()

	viper.BindPFlag(flagForGrpcPort, flag.Lookup(flagForGrpcPort))

	viper.AutomaticEnv()

	c.GrpcPort = viper.GetInt(flagForGrpcPort)

	c.LoggingConfig = logging.LoadConfig()

	return c
}
