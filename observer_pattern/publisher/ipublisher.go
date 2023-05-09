package publisher

import "design-pattern-golang-example/observer_pattern/subscriber"

type IPublisher interface {
	Subscribe(s subscriber.ISubscriber)
	UnSubscribe(s subscriber.ISubscriber)
	NotifySubscribes()
}
