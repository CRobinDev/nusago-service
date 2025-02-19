package config

import (
	"os"

	"github.com/CRobinDev/nusago-service/internal/entity"
	"github.com/CRobinDev/nusago-service/pkg/response"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func loadDSN() string {
// 	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_NAME"),
// 		os.Getenv("DB_PORT"),
// 	)
// }

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
