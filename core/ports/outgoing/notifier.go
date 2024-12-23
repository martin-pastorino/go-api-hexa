package outgoing

type Notifier interface {
	SendWelcomeEmail(email string) error
}
