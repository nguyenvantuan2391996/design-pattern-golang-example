package main

import (
	"time"

	"design-pattern-golang-example/observer_pattern/publisher/concrete_publisher"
	"design-pattern-golang-example/observer_pattern/subscriber"
	"design-pattern-golang-example/observer_pattern/subscriber/concrete_subscriber"
)

func main() {
	// init publisher
	appleStore := concrete_publisher.NewAppleStore([]subscriber.ISubscriber{}, "apple store")

	// init observer
	subscriberOne := concrete_subscriber.NewCustomer("subscribe-one@gmail.com")
	subscriberTwo := concrete_subscriber.NewCustomer("subscribe-two@gmail.com")

	// the observer subscribes publisher
	appleStore.Subscribe(subscriberOne)
	appleStore.Subscribe(subscriberTwo)

	// send notification
	time.Sleep(3 * time.Second)
	appleStore.NotifySubscribes()

	// the observer unsubscribes publisher
	appleStore.UnSubscribe(subscriberTwo)

	// send notification
	time.Sleep(3 * time.Second)
	appleStore.NotifySubscribes()
}
