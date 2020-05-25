package metrics

import (
	"net/http"

	"github.com/clintjedwards/comet/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/zerolog/log"
)

// InitPrometheusService starts a long running http prometheus endpoint
func InitPrometheusService(config *config.Config) {

	log.Info().Str("url", config.Metrics.Endpoint).Msg("starting metrics http service")

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal().Err(http.ListenAndServe(config.Metrics.Endpoint, nil)).Msg("metrics server exited abnormally")
}
