package main

import (
	"context"
	"errors"
	"fmt"
	"test-task/internal/config"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Subscription struct {
	SubscriptionID uuid.UUID `json:"subscription_id" db:"subscription_id"`
	ServiceName    string    `json:"service_name" db:"service_name"`
	Price          int       `json:"price" db:"price"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	StartDate      string    `json:"start_date" db:"start_date"`
	EndDate        *string   `json:"end_date,omitempty" db:"end_date"`
}

// TODO: POST /api/v1/subscriptions -> create

// TODO: GET /api/v1/subscriptions/{subscription_id} -> read

// TODO: PATCH /api/v1/subscriptions/{subscription_id} -> update

// TODO: DELETE /api/v1/subscriptions/{subscription_id} -> delete

// TODO: GET /api/v1/subscriptions?date_from=...&date_to=...&user_id=...&limit=...&offset=...&service_name=... -> list

// TODO: GET /api/v1/subscriptions/total?date_from=07-2025&date_to=08-2025&user_id=...&service_name=...

func DateFormatting(date time.Time) string {
	year, month, _ := date.Date()
	return fmt.Sprintf("%02d-%04d", int(month), year)
}

func ConnectToDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.PostgresConfig.Host,
		cfg.PostgresConfig.User,
		cfg.PostgresConfig.Password,
		cfg.PostgresConfig.Dbname,
		cfg.PostgresConfig.Port,
		cfg.PostgresConfig.SSLMode,
		cfg.PostgresConfig.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) CreateSubscription(ctx context.Context, sub *Subscription) error {
	if err := r.db.WithContext(ctx).Create(sub).Error; err != nil {
		return fmt.Errorf("create subscription: %w", err)
	}
	return nil
}

func (r *SubscriptionRepository) ReadSubscriptions(ctx context.Context, subscriptionID uuid.UUID) (*Subscription, error) {
	var sub *Subscription
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

func main() {
	s := &Subscription{
		SubscriptionID: uuid.New(),
		ServiceName:    "yandex_plus",
		Price:          100,
		UserID:         uuid.New(),
		StartDate:      DateFormatting(time.Now()),
		EndDate:        new(DateFormatting(time.Now())),
	}

	fmt.Printf("SubscriptionID: %s\nServiceName: %s\nPrice: %d\nUserID: %s\nStartDate: %s\nEndDate: %s\n",
		s.SubscriptionID,
		s.ServiceName,
		s.Price,
		s.UserID,
		s.StartDate,
		*s.EndDate,
	)

	fmt.Println()

	cfg, err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := ConnectToDatabase(cfg)
	if err != nil {
		panic(err)
	}
	_ = db

	err = db.AutoMigrate(&Subscription{})
	if err != nil {
		panic(err)
	}

	database := NewSubscriptionRepository(db)
	err = database.CreateSubscription(context.Background(), s)
	if err != nil {
		panic(err)
	}

	sub, err := database.ReadSubscriptions(context.Background(), uuid.MustParse("edfae7fc-bb3f-414c-8ddd-cb105d04049d"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("SubscriptionID: %s\nServiceName: %s\nPrice: %d\nUserID: %s\nStartDate: %s\nEndDate: %s\n",
		sub.SubscriptionID,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate,
		*sub.EndDate,
	)
}
