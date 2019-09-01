package service

import (
	"github.com/google/uuid"
	"log"
	"metrics-collector/dao"
	"metrics-collector/model"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "service: ", log.Lshortfile)
}

type Service struct {
	dao dao.Dao
}

func New(dao dao.Dao) *Service {
	logger.Println("New service instance was initialized")
	return &Service{
		dao: dao,
	}
}

func (s *Service) SubmitMetric(metric model.Metric) {
	id := uuid.New().String()
	logger.Printf("Store metric with id '%s' in storage", id)
	s.dao.StoreMetric(id, metric)
}