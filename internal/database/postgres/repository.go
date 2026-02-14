package postgres

import (
	"context"
	"errors"
	"fmt"
	"test-task/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) CreateSubscription(ctx context.Context, sub *model.Subscription) error {
	if err := r.db.WithContext(ctx).Create(sub).Error; err != nil {
		return fmt.Errorf("create subscription: %w", err)
	}
	return nil
}

func (r *SubscriptionRepository) ReadSubscriptions(ctx context.Context, subscriptionID uuid.UUID) (*model.Subscription, error) {
	var sub *model.Subscription
	if err := r.db.WithContext(ctx).First(&sub, subscriptionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("read subscription: %w", err)
		}
		return nil, fmt.Errorf("read subscription: %w", err)
	}
	return sub, nil
}

func UpdateSubscription() {}

func DeleteSubscription() {}

func ListSubscriptions() {}
