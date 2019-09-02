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
	logger = log.New(os.Stdout, "dao: ", log.Lshortfile)
)

type (
	Metrics map[string]model.Metric

	InMemoryStorage struct {
		metrics  Metrics
		lock     sync.RWMutex
		filename string
	}
)

// Creates a new storage instance.
// Starts ticker to flush storage repeatedly
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

// Stores received metric in storage
func (s *InMemoryStorage) StoreMetric(id string, metric model.Metric) {
	s.lock.Lock()
	s.metrics[id] = metric
	s.lock.Unlock()
	logger.Printf("Metric with id '%s' was stored successfully", id)
}

// Flushes storage if it's not empty
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

// Returns all metrics from storage in a new collection
func (s *InMemoryStorage) getAll() Metrics {
	logger.Println("Get all metrics from storage")
	metrics := make(Metrics, len(s.metrics))
	for k, v := range s.metrics {
		metrics[k] = v
	}

	logger.Printf("Return '%d' metrics from storage", len(metrics))
	return metrics
}

// Writes received metrics to file.
// Creates a new file if it doesn't exist
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
			if e := file.Close(); e != nil {
				logger.Printf("Failed to close the file: '%s'", e.Error())
			}
			return err
		}
	}
	return file.Close()
}
