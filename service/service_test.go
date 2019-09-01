package service

import (
	"github.com/golang/mock/gomock"
	"metrics-collector/dao/mock_dao"
	"metrics-collector/model"
	"testing"
	"time"
)

var (
	mockCtrl *gomock.Controller
	mockDao  *mock_dao.MockDao
	metric   = model.Metric{
		EventType: "click",
		UserAgent: "safari",
		Timestamp: time.Now().UnixNano(),
	}
)

func testInitializer(t *testing.T) {
	mockCtrl = gomock.NewController(t)
	mockDao = mock_dao.NewMockDao(mockCtrl)
}

func TestService_SubmitMetric(t *testing.T) {
	testInitializer(t)
	defer mockCtrl.Finish()

	gomock.InOrder(
		mockDao.
			EXPECT().
			StoreMetric(gomock.Any(), gomock.Eq(metric)).
			Times(1),
	)

	service := New(mockDao)
	service.SubmitMetric(metric)
}
