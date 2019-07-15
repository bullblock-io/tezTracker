package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config stores the required configuration.
type Config struct {
	Port                int
	SqlConnectionString string `envconfig:"SQL_CONNECTION_STRING"`
	LogLevel            string `envconfig:"LOG_LEVEL"`
}

// Parse initializes the configuration from .env file or from environment.
func Parse() Config {
	var s Config
	_ = godotenv.Load()

	err := envconfig.Process("teztracker", &s) // "TEZTRACKER_" is prefix for every env config variable
	if err != nil {
		log.Fatal(err.Error())
	}

	log.SetLevel(parseLogLevel(s.LogLevel))

	return s
}

func parseLogLevel(lvl string) log.Level {
	level, err := log.ParseLevel(lvl)
	if err != nil {
		return log.InfoLevel
	}
	return level
}

// DbLogger is a simple log wrapper for use with gorm and logrus.
type DbLogger struct{}

// Print directs the log ouput to trace level.
func (*DbLogger) Print(args ...interface{}) {
	log.Traceln(args...)
}
