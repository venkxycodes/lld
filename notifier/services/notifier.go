package services

import "notifier/domain"

type notifier struct {
	queue Queuer
}

type Notifier interface {
	SendSms(queue *domain.Queue, notification *domain.Notification) error
	SendEmail(queue *domain.Queue, notification *domain.Notification) error
}

func NewNotifier(queue Queuer) Notifier {
	return &notifier{
		queue: queue,
	}
}

func (n notifier) SendSms(queue *domain.Queue, notification *domain.Notification) error {
	err := n.queue.EnqueueJob(queue, notification)
	if err != nil {
		return err
	}
	notification.Status = domain.Sent
	return nil
}

func (n notifier) SendEmail(queue *domain.Queue, notification *domain.Notification) error {
	err := n.queue.EnqueueJob(queue, notification)
	if err != nil {
		return err
	}
	notification.Status = domain.Sent
	return nil
}
