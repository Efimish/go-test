package notifications

import (
	"fmt"
	"time"
)

type NotificationService struct {
	notifications map[uint][]Notification
}

func NewNotificationService(publicURL string) NotificationService {
	image1 := fmt.Sprintf("%s/static/anon1.webp", publicURL)
	image2 := fmt.Sprintf("%s/static/anon2.jpg", publicURL)
	image3 := fmt.Sprintf("%s/static/anon3.webp", publicURL)

	return NotificationService{
		notifications: map[uint][]Notification{
			1: {
				{
					Date:  time.Now().Add(-time.Minute * 30),
					Title: "Message from John",
					Body:  "Hello",
					Image: &image1,
				},
				{
					Date:  time.Now().Add(-time.Minute * 20),
					Title: "Message from Robert",
					Body:  "Hey bro",
					Image: &image2,
				},
				{
					Date:  time.Now().Add(-time.Minute * 10),
					Title: "Message from Kyle",
					Body:  "Damn",
					Image: &image3,
				},
			},
			2: {
				{
					Date:  time.Now().Add(-time.Minute * 60),
					Title: "Message from Ashley",
					Body:  "See you there!",
				},
				{
					Date:  time.Now().Add(-time.Minute * 10),
					Title: "Message from Ashley",
					Body:  "Where are you?",
				},
			},
			3: {
				{
					Date:  time.Now(),
					Title: "Message from Boss",
					Body:  "Come here right now!",
				},
			},
		},
	}
}

func (s NotificationService) ListByUserID(userID uint) []Notification {
	return s.notifications[userID]
}
