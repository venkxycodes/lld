package main

import (
	"fmt"
	"notifier/domain"
	"notifier/services"
	"os"
	"time"
)

func main() {
	smsQueue := domain.NewQueue("SMS Queue", domain.SMS, 20)
	emailQueue := domain.NewQueue("Email Queue", domain.Email, 20)
	deadLetterQueue := domain.NewQueue("Dead Queue", domain.DeadLetter, 40)

	queuer := services.NewQueuer()
	notifier, err := services.NewNotifier(queuer, map[domain.QueueType]*domain.Queue{
		domain.SMS:   smsQueue,
		domain.Email: emailQueue,
	})
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error(), "err-notifier-missing-one-of-the-message-queues"))
		os.Exit(0)
	}
	notification1 := &domain.Notification{
		Id:             "uuid-1",
		Message:        "Hey, this is from SMS",
		CreatedAt:      time.Now().UnixMilli(),
		ScheduledAfter: 0,
		UpdatedAt:      time.Now().UnixMilli(),
	}
	notification2 := &domain.Notification{
		Id:             "uuid-2",
		Message:        "Hey, this is from Email",
		CreatedAt:      time.Now().UnixMilli(),
		ScheduledAfter: 0,
		UpdatedAt:      time.Now().UnixMilli(),
	}

	smsErr := notifier.SendSms(notification1)
	if smsErr != nil {
		enqueueErr := queuer.EnqueueJob(deadLetterQueue, notification1)
		if enqueueErr != nil {
			return
		}
		fmt.Println(smsErr)
	}

	emailErr := notifier.SendEmail(notification2)
	if emailErr != nil {
		fmt.Println(emailErr)
		enqueueErr := queuer.EnqueueJob(deadLetterQueue, notification2)
		if enqueueErr != nil {
			return
		}
	}

	job, _ := queuer.DequeueJob(emailQueue)
	fmt.Println(job)
	job2, _ := queuer.DequeueJob(smsQueue)
	fmt.Println(job2)
}
