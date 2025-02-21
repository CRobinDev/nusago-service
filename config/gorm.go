package config

import (
	"os"

	"github.com/CRobinDev/nusago-service/internal/entity"
	"github.com/CRobinDev/nusago-service/pkg/response"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, &response.ErrConnectDatabase
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		return &response.ErrMigrateDatabase
	}
	return nil
}
