package domain

type QueueType string

const (
	SMS        QueueType = "SMS"
	Email      QueueType = "Email"
	DeadLetter QueueType = "DeadLetter"
)

type Queue struct {
	Name    string
	Type    QueueType
	MaxSize int64
	Values  []Notification
}

func NewQueue(name string, qType QueueType, maxSize int64) *Queue {
	q := &Queue{}
	q.Name = name
	q.Type = qType
	q.MaxSize = maxSize
	q.Values = make([]Notification, 0)
	return q
}
