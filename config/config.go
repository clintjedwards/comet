package config

import "github.com/kelseyhightower/envconfig"

// BoltDBConfig represents a on-disk key/value store
// https://github.com/boltdb/bolt
type BoltDBConfig struct {
	// file path for database file
	Path string `envconfig:"database_path" default:"/tmp/comet.db"`
}

// BackendConfig defines settings for comet's backend assets
type BackendConfig struct {
	// The path where built plugin binaries should stay
	PluginDirectoryPath string `envconfig:"plugin_directory_path" default:"/tmp/comet-dev"`
}

// CometConfig defines config settings for the comet service
type CometConfig struct {
	// the length of all randomly generated ids
	IDLength int    `envconfig:"id_length" default:"5"`
	HTTPURL  string `envconfig:"http_url" default:"localhost:8080"`
	GRPCURL  string `envconfig:"grpc_url" default:"localhost:8081"`
	// the total amount of comets a single user is allowed to checkout
	// A value of -1 is unlimited
	UserLimit int `envconfig:"user_limit" default:"-1"`
	// max time that a comet can be requested
	// accepts humanized time strings: 1s = 1 second / 4d = 4 days / 3w = 3 weeks
	MaxDuration string `envconfig:"max_duration" default:"3d"`
	// number of seconds prune thread waits before scanning for comets that have
	// passed their intended duration
	PruneInterval int `envconfig:"prune_interval" default:""`
}

// CommandLineConfig represents configuration for cli application
type CommandLineConfig struct {
	Token string `envconfig:"cli_token" default:"test"`
}

// DatabaseConfig defines config settings for comet database
type DatabaseConfig struct {
	// The database engine used by the backend
	// possible values are: boltdb
	Engine string `envconfig:"database_engine" default:"boltdb"`
	BoltDB *BoltDBConfig
}

// MetricsConfig represents configuration for the metrics endpoint
type MetricsConfig struct {
	Endpoint string `envconfig:"metrics_endpoint" default:"localhost:8082"`
}

// Config represents overall configuration objects of the application
type Config struct {
	Debug       bool   `envconfig:"debug" default:"false"`
	TLSCertPath string `envconfig:"tls_cert_path" default:"./localhost.crt"`
	TLSKeyPath  string `envconfig:"tls_key_path" default:"./localhost.key"`
	Comet       *CometConfig
	CommandLine *CommandLineConfig
	Database    *DatabaseConfig
	Metrics     *MetricsConfig
	Backend     *BackendConfig
}

// FromEnv parses environment variables into the config object based on envconfig name
func FromEnv() (*Config, error) {
	var config Config
	err := envconfig.Process("comet", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
