package instance

import (
	"log/slog"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/ncostamagna/go-monitoring/app/internal/product"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const METHOD = "method"

func NewProductService(logger *slog.Logger) product.Service {

	fieldKeys := []string{METHOD}
	repository := product.NewRepo(logger)
	service := product.NewService(logger, repository)
	return product.NewInstrumenting(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "product_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "product_service",
			Name:      "request_latency_microseconds_summary",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: "api",
			Subsystem: "product_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		service)

}
