// Comet is a temporary machine procurement and management system
package main

import (
	"github.com/clintjedwards/comet/cmd"
	"github.com/clintjedwards/comet/config"
	"github.com/rs/zerolog/log"
)

func main() {

	conf, err := config.FromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("could not load env config")
	}

	setupLogging(conf.LogLevel, conf.Debug)

	cmd.RootCmd.Execute()
}
