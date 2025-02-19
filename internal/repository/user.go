package repository

import (
	"errors"
	"time"

	"github.com/CRobinDev/nusago-service/internal/entity"
	"github.com/CRobinDev/nusago-service/pkg/response"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *entity.User) error
	FindByID(id uuid.UUID) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Update(user *entity.User) error
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *entity.User) error {
	tr := ur.db.Begin()
	
	if err := tr.Create(&user).Error; err != nil {
		tr.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return &response.ErrUserAlreadyExists
		}
		return err
	}

	return tr.Commit().Error
}

func (ur *userRepository) FindByID(id uuid.UUID) (entity.User, error) {
	var user entity.User
	err := ur.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, &response.ErrUserNotFound
	}

	return user, err
}

func (ur *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.User{}, &response.ErrUserNotFound
	}

	return user, err
}

func (ur *userRepository) Update(user *entity.User) error {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	
	result := ur.db.Model(&user).Where("id", user.ID).Updates(map[string]interface{}{
		"fullname":    user.Fullname,
		"institution": user.Institution,
		"description": user.Description,
		"updated_at":  updatedAt,
	})

	if result.RowsAffected == 0 {
		return &response.ErrUserNotFound
	}

	return result.Error
}

func (ur *userRepository) Delete(id uuid.UUID) error {
	var user entity.User
	result := ur.db.Delete(&user, id)
	if result.RowsAffected == 0 {
		return &response.ErrUserNotFound
	}

	return result.Error
}
