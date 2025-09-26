package outgoing

import (
	"api/core/domain"
	"context"
)


type PurchaseRepository interface {
	Save(ctx context.Context, purchase domain.Purchase) (string, error)
	GetPurchase(ctx context.Context, id string) (domain.Purchase, error)
	Search(ctx context.Context, id string) (domain.Purchase, error)
}