package main

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

}
