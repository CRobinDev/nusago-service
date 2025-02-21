package jwt

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/CRobinDev/nusago-service/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type IJWT interface {
	CreateToken(user *entity.User) (string, error)
	DecodeToken(tokenString string) (uuid.UUID, error)
}

type token struct {
	SecretKey          string
	ExpiredTime        time.Duration
	RefreshExpiredTime time.Duration
}

type Claims struct {
	UserID      uuid.UUID `json:"user_id"`
	Username    string    `json:"username"`
	Fullname    string    `json:"fullname"`
	jwt.RegisteredClaims
}

func Init() IJWT {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	duration := os.Getenv("JWT_EXPIRED_TIME")

	expTime, err := strconv.Atoi(duration)
	if err != nil {
		log.Printf("failed parse duration")
	}

	return &token{
		SecretKey:   secretKey,
		ExpiredTime: time.Duration(expTime) * time.Minute,
	}
}

func (t *token) CreateToken(user *entity.User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Fullname: user.Fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.ExpiredTime)),
		},
	}
	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedJWT, err := unsignedJWT.SignedString([]byte(t.SecretKey))

	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func (t *token) DecodeToken(tokenString string) (uuid.UUID, error) {
	var claims Claims
	var userID uuid.UUID

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return []byte(t.SecretKey), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, jwt.ErrSignatureInvalid
	}

	userID = claims.UserID

	return userID, nil
}

func GetUser(c *fiber.Ctx) (uuid.UUID, error) {
	claims, ok := c.Locals("claims").(uuid.UUID)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Can't retrieve claims")
	}

	userID := claims

	return userID, nil
}
