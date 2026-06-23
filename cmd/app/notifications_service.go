package main

import (
	"fmt"
	"time"
)

var image1 = fmt.Sprintf("http://%s/static/anon1.webp", Config.Host)
var image2 = fmt.Sprintf("http://%s/static/anon2.jpg", Config.Host)
var image3 = fmt.Sprintf("http://%s/static/anon3.webp", Config.Host)
var notifications = map[uint][]Notification{
	1: []Notification{
		Notification{
			Date:  time.Now().Add(-time.Minute * 30),
			Title: "Message from John",
			Body:  "Hello",
			Image: &image1,
		},
		Notification{
			Date:  time.Now().Add(-time.Minute * 20),
			Title: "Message from Robert",
			Body:  "Hey bro",
			Image: &image2,
		},
		Notification{
			Date:  time.Now().Add(-time.Minute * 10),
			Title: "Message from Kyle",
			Body:  "Damn",
			Image: &image3,
		},
	},
	2: []Notification{
		Notification{
			Date:  time.Now().Add(-time.Minute * 60),
			Title: "Message from Ashley",
			Body:  "See you there!",
		},
		Notification{
			Date:  time.Now().Add(-time.Minute * 10),
			Title: "Message from Ashley",
			Body:  "Where are you?",
		},
	},
	3: []Notification{
		Notification{
			Date:  time.Now(),
			Title: "Message from Boss",
			Body:  "Come here right now!",
		},
	},
}
