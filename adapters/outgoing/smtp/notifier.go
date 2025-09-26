package smtp

import (
	"api/core/ports/outgoing"
	"context"
	"fmt"
)

type Notifier struct {}

func NewNotifier() *Notifier {
	return &Notifier{}
}

// Provider for Notifier
func NewNotifierProvider() outgoing.Notifier {
    return NewNotifier()
}

func (n *Notifier) SendWelcomeEmail(ctx context.Context, email string) error {
	// Send welcome email
	fmt.Println("Welcome email sent to", email)
	return nil
}