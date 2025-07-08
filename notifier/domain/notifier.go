package domain

type Status string

const (
	Sent     Status = "Sent"
	Received Status = "Received"
	Seen     Status = "Seen"
	Failed   Status = "Failed"
)

type Notification struct {
	Id             string
	Message        string
	CreatedAt      int64 // epoch
	ScheduledAfter int64
	UpdatedAt      int64
	Status         Status
}
