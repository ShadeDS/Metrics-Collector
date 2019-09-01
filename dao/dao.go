package dao

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"metrics-collector/model"
	"os"
	"sync"
	"time"
)

//go:generate ${GOPATH}/bin/mockgen -source=dao.go -destination=mock_dao/mock.go
type Dao interface {
	StoreMetric(id string, metric model.Metric)
}

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "dao: ", log.Lshortfile)
}

type (
	Metrics map[string]model.Metric

	InMemoryStorage struct {
		metrics  Metrics
		lock     sync.RWMutex
		filename string
	}
)

func New() Dao {
	storage := &InMemoryStorage{
		metrics:  make(Metrics),
		filename: uuid.New().String(),
	}
	logger.Printf("New storage instance was initialized with target filename '%s'", storage.filename)

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				storage.flush()
			}
		}
	}()

	return storage
}

func (s *InMemoryStorage) StoreMetric(id string, metric model.Metric) {
	s.lock.Lock()
	s.metrics[id] = metric
	s.lock.Unlock()
	logger.Printf("Metric with id '%s' was stored successfully", id)
}

func (s *InMemoryStorage) flush() {
	logger.Println("Flush metrics from in-memory storage")
	s.lock.Lock()
	if len(s.metrics) > 0 {
		metrics := s.getAll()
		s.metrics = make(Metrics)
		s.lock.Unlock()

		err := s.save(metrics)
		if err != nil {
			logger.Printf("Error occurred while saving metrics: '%s'", err.Error())
		} else {
			logger.Println("Metrics were saved successfully")
		}
	} else {
		logger.Println("Storage is empty, nothing to do here")
		s.lock.Unlock()
	}
}

func (s *InMemoryStorage) getAll() Metrics {
	logger.Println("Get all metrics from storage")
	metrics := make(Metrics, len(s.metrics))
	for k, v := range s.metrics {
		metrics[k] = v
	}

	logger.Printf("Return '%d' metrics from storage", len(metrics))
	return metrics
}

func (s *InMemoryStorage) save(metrics Metrics) error {
	logger.Printf("Save '%d' metrics to file '%s'", len(metrics), s.filename)
	file, err := os.OpenFile(s.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logger.Println("Encode metrics to file")
	for _, v := range metrics {
		b, _ := json.Marshal(v)
		_, err = file.WriteString(string(b) + "\n")
		if err != nil {
			file.Close()
			return err
		}
	}
	return file.Close()
}
