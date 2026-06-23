package main

import "time"

type Notification struct {
	Date  time.Time `json:"date"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Image *string   `json:"image,omitempty"`
}
