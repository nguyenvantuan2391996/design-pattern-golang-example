package concrete_subscriber

import (
	"fmt"

	"design-pattern-golang-example/observer_pattern/subscriber"
)

type Customer struct {
	Email string
}

func NewCustomer(email string) subscriber.ISubscriber {
	return &Customer{Email: email}
}

func (c *Customer) SendEmail(notification string) {
	content := fmt.Sprintf(`Sending the notification for %s with content "%s"`, c.GetEmail(), notification)
	fmt.Println(content)
}

func (c *Customer) GetEmail() string {
	return c.Email
}
