package outgoing

import "context"

type Notifier interface {
	SendWelcomeEmail(ctx context.Context, email string) error
}
