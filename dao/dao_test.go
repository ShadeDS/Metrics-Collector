package dao

import (
	"metrics-collector/model"
	"testing"
)

var (
	id     = "metricId"
	metric = model.Metric{
		EventType: "click",
	}
	metrics = make(Metrics)
)

func TestInMemoryStorage_StoreMetric(t *testing.T) {
	storage := InMemoryStorage{
		metrics:  metrics,
		filename: "test",
	}

	storage.StoreMetric(id, metric)
	if _, ok := metrics[id]; !ok {
		t.Fatal("Metrics do not contain stored metric")
	}
}
