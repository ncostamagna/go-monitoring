package product

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/ncostamagna/go-monitoring/app/internal/domain"
)

type (
	instrumenting struct {
		requestCount          metrics.Counter
		requestLatency        metrics.Histogram
		requestLatencySummary metrics.Histogram
		s                     Service
	}

	Instrumenting interface {
		Service
	}
)

func NewInstrumenting(requestCount metrics.Counter, requestLatencySummary metrics.Histogram, requestLatency metrics.Histogram, s Service) Instrumenting {
	return &instrumenting{
		requestCount:          requestCount,
		requestLatencySummary: requestLatencySummary,
		requestLatency:        requestLatency,
		s:                     s,
	}
}

func (i *instrumenting) Store(ctx context.Context, name, description string, price float64) (*domain.Product, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Store").Add(1)
		i.requestLatencySummary.With("method", "Store").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Store").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Store(ctx, name, description, price)
}

func (i *instrumenting) Get(ctx context.Context, id int) (*domain.Product, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Get").Add(1)
		i.requestLatencySummary.With("method", "Get").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Get").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Get(ctx, id)
}

func (i *instrumenting) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Product, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "GetAll").Add(1)
		i.requestLatencySummary.With("method", "GetAll").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.GetAll(ctx, filters, offset, limit)
}

func (i *instrumenting) Delete(ctx context.Context, id int) error {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Delete").Add(1)
		i.requestLatencySummary.With("method", "Delete").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Delete(ctx, id)
}

func (i *instrumenting) Update(ctx context.Context, id int, name, description *string, price *float64) error {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Update").Add(1)
		i.requestLatencySummary.With("method", "Update").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Update").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Update(ctx, id, name, description, price)
}

func (i *instrumenting) Count(ctx context.Context, filters Filters) (int, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Count").Add(1)
		i.requestLatencySummary.With("method", "Count").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Count").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Count(ctx, filters)
}
