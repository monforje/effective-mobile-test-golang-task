package app

import (
	"test-task/internal/config"
	"test-task/internal/database/postgres"
	"test-task/internal/model"
)

// TODO: POST /api/v1/subscriptions -> create

// TODO: GET /api/v1/subscriptions/{subscription_id} -> read

// TODO: PATCH /api/v1/subscriptions/{subscription_id} -> update

// TODO: DELETE /api/v1/subscriptions/{subscription_id} -> delete

// TODO: GET /api/v1/subscriptions?date_from=...&date_to=...&user_id=...&limit=...&offset=...&service_name=... -> list

// TODO: GET /api/v1/subscriptions/total?date_from=07-2025&date_to=08-2025&user_id=...&service_name=...

type App struct{}

func New() (*App, error) {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		return nil, err
	}

	db, err := postgres.ConnectToDatabase(cfg)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Subscription{})
	if err != nil {
		return nil, err
	}

	subscriptionRepository := postgres.NewSubscriptionRepository(db)
	_ = subscriptionRepository

	return &App{}, nil
}
