package notifications

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	notifications map[uint][]Notification
}

func NewService(publicURL string) Service {
	return Service{
		notifications: map[uint][]Notification{
			1: {
				{
					Id:       uuid.New(),
					Date:     time.Now().Add(-time.Minute * 30),
					Title:    new("Message from John"),
					Icon:     new(fmt.Sprintf("%s/static/anon1.webp", publicURL)),
					Priority: PriorityHigh,
					Body:     "Hello",
				},
				{
					Id:       uuid.New(),
					Date:     time.Now().Add(-time.Minute * 20),
					Title:    new("Message from Robert"),
					Icon:     new(fmt.Sprintf("%s/static/anon2.jpg", publicURL)),
					Priority: PriorityHigh,
					Body:     "Hey bro",
				},
				{
					Id:       uuid.New(),
					Date:     time.Now().Add(-time.Minute * 10),
					Title:    new("Message from Kyle"),
					Icon:     new(fmt.Sprintf("%s/static/anon3.webp", publicURL)),
					Priority: PriorityHigh,
					Body:     "Damn",
				},
			},
			2: {
				{
					Id:       uuid.New(),
					Date:     time.Now().Add(-time.Minute * 60),
					Title:    new("Message from Ashley"),
					Priority: PriorityHigh,
					Body:     "See you there!",
				},
				{
					Id:       uuid.New(),
					Date:     time.Now().Add(-time.Minute * 10),
					Title:    new("Message from Ashley"),
					Priority: PriorityHigh,
					Body:     "Where are you?",
				},
			},
			3: {
				{
					Id:       uuid.New(),
					Date:     time.Now(),
					Title:    new("Message from Boss"),
					Priority: PriorityHigh,
					Body:     "Come here right now!",
				},
			},
		},
	}
}

func (s Service) AmountByUserId(userId uint) int {
	return len(s.notifications[userId])
}

func (s Service) ListByUserId(userId uint) []Notification {
	return s.notifications[userId]
}

func (s Service) DeleteByIds(userId uint, ids uuid.UUIDs) {
	s.notifications[userId] = slices.DeleteFunc(s.notifications[userId], func(n Notification) bool {
		return slices.Contains(ids, n.Id)
	})
}
