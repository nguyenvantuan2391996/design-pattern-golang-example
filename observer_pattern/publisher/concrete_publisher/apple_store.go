package concrete_publisher

import (
	"fmt"
	"time"

	"design-pattern-golang-example/observer_pattern/publisher"
	"design-pattern-golang-example/observer_pattern/subscriber"
)

type AppleStore struct {
	Subscribers []subscriber.ISubscriber
	Name        string
}

func NewAppleStore(subscribers []subscriber.ISubscriber, name string) publisher.IPublisher {
	return &AppleStore{
		Subscribers: subscribers,
		Name:        name,
	}
}

func (as *AppleStore) Subscribe(s subscriber.ISubscriber) {
	as.Subscribers = append(as.Subscribers, s)
	time.Sleep(1 * time.Second)
	fmt.Println(s.GetEmail() + " was subscribed...")
}

func (as *AppleStore) UnSubscribe(s subscriber.ISubscriber) {
	subscribersUpdate := make([]subscriber.ISubscriber, 0)
	for _, sub := range as.Subscribers {
		if sub.GetEmail() != s.GetEmail() {
			subscribersUpdate = append(subscribersUpdate, sub)
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Println(s.GetEmail() + " was unsubscribed...")

	as.Subscribers = subscribersUpdate
}

func (as *AppleStore) NotifySubscribes() {
	for _, sub := range as.Subscribers {
		sub.SendEmail("We have had a new model of the iPhone")
		time.Sleep(1 * time.Second)
	}
}
