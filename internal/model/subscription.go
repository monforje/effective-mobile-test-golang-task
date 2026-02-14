package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	SubscriptionID uuid.UUID `json:"subscription_id" db:"subscription_id"`
	ServiceName    string    `json:"service_name" db:"service_name"`
	Price          int       `json:"price" db:"price"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	StartDate      string    `json:"start_date" db:"start_date"`
	EndDate        *string   `json:"end_date,omitempty" db:"end_date"`
}

func DateFormatting(date time.Time) string {
	year, month, _ := date.Date()
	return fmt.Sprintf("%02d-%04d", int(month), year)
}
