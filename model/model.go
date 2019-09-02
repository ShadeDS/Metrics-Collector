package model

import "errors"

type Metric struct {
	EventType string
	UserAgent string
	Timestamp int64
}

func (m *Metric) Validate() error {
	if m.EventType == "" {
		return errors.New("metric is invalid. event type is empty")
	} else if m.UserAgent == "" {
		return errors.New("metric is invalid. user agent is empty")
	} else if m.Timestamp == 0 {
		return errors.New("metric is invalid. timestamp is not specified")
	} else {
		return nil
	}
}
