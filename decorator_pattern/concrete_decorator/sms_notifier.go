package concrete_decorator

import (
	"fmt"
	"time"

	"design-pattern-golang-example/decorator_pattern/component"
)

type SmsNotifier struct {
	Notifier component.INotifier
}

func (s *SmsNotifier) SendNotification(msg string) {
	s.Notifier.SendNotification(msg)
	fmt.Println(fmt.Sprintf("Send the notification to SMS with message %s", msg))
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
