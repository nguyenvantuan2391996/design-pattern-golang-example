package subscriber

// ISubscriber this is the observer
type ISubscriber interface {
	SendEmail(notification string)
	GetEmail() string
}
