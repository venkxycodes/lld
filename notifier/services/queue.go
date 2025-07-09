package services

import (
	"fmt"
	"notifier/domain"
)

type queuer struct{}

type Queuer interface {
	EnqueueJob(queue *domain.Queue, notification *domain.Notification) error
	DequeueJob(queue *domain.Queue) (*domain.Notification, error)
}

func NewQueuer() Queuer {
	return &queuer{}
}

func (q queuer) EnqueueJob(queue *domain.Queue, notification *domain.Notification) error {
	if queue.MaxSize == int64(len(queue.Values)) {
		return fmt.Errorf("max size of %d is reached", queue.MaxSize)
	}
	queue.Values = append(queue.Values, *notification)
	return nil
}

func (q queuer) DequeueJob(queue *domain.Queue) (*domain.Notification, error) {
	if len(queue.Values) == 0 {
		return nil, fmt.Errorf(`queuer "%s" is empty`, queue.Name)
	}
	notification := &queue.Values[0]
	queue.Values = queue.Values[1:]
	return notification, nil
}
