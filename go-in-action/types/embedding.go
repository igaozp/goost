package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // type embedding
	level string
}

// if the following code is commented out,
// the notify operation in the main method will invoke the user's notify method
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@gmail.com",
		},
		level: "super",
	}

	ad.user.notify()

	ad.notify()

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
