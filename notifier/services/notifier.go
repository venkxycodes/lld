package services

import (
	"errors"
	"notifier/domain"
)

type notifier struct {
	smsQueue   *domain.Queue
	emailQueue *domain.Queue
	queuer     Queuer
}

type Notifier interface {
	SendSms(notification *domain.Notification) error
	SendEmail(notification *domain.Notification) error
}

func NewNotifier(queuer Queuer, queues map[domain.QueueType]*domain.Queue) (Notifier, error) {
	var smsQueue *domain.Queue
	var emailQueue *domain.Queue
	if smsQueue, ok = queues[domain.SMS]; !ok {
		return nil, errors.New("no sms queue")
	}
	if emailQueue, ok = queues[domain.Email]; !ok {
		return nil, errors.New("no email queue")
	}
	return &notifier{
		queuer:     queuer,
		smsQueue:   smsQueue,
		emailQueue: emailQueue,
	}, nil
}

func (n notifier) SendSms(notification *domain.Notification) error {
	err := n.queuer.EnqueueJob(n.smsQueue, notification)
	if err != nil {
		return err
	}
	notification.Status = domain.Sent
	return nil
}

func (n notifier) SendEmail(notification *domain.Notification) error {
	err := n.queuer.EnqueueJob(n.emailQueue, notification)
	if err != nil {
		return err
	}
	notification.Status = domain.Sent
	return nil
}
