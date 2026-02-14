package postgres

import (
	"fmt"
	"test-task/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
