package notifications

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type NotificationPriority string

const (
	PriorityLow      NotificationPriority = "low"
	PriorityMedium   NotificationPriority = "medium"
	PriorityHigh     NotificationPriority = "high"
	PriorityCritical NotificationPriority = "critical"
)

func (p NotificationPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

func (p *NotificationPriority) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch NotificationPriority(s) {
	case PriorityLow, PriorityMedium, PriorityHigh, PriorityCritical:
		*p = NotificationPriority(s)
		return nil
	default:
		return fmt.Errorf("invalid priority: %s", s)
	}
}

type Notification struct {
	Id       uuid.UUID            `json:"id"`
	Date     time.Time            `json:"date"`
	Title    *string              `json:"title,omitempty"`
	Icon     *string              `json:"icon,omitempty"`
	Priority NotificationPriority `json:"priority"`
	Body     string               `json:"body"`
}
