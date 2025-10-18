package authservice

import (
	"fmt"
	"strings"
	entity "suggestApp/enity"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	SignKey               string
	AccessExpiretionTime  time.Duration
	RefreshExpiretionTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

type Service struct {
	config Config
}

func New(cfg Config) Service {

	return Service{
		config: cfg,
	}

}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.AccessSubject, s.config.AccessExpiretionTime)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.RefreshSubject, s.config.RefreshExpiretionTime)
}

func (s Service) ParseToken(BearerToken string) (*Claims, error) {

	tokenStr := strings.Replace(BearerToken, "bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})

	if token == nil {

		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("userID %v ExpiresAt %v\n", claims.UserID, claims.RegisteredClaims.ExpiresAt)

		return claims, nil
	} else {

		return nil, err
	}

}

func (s Service) createToken(userID uint, subject string, expiteAt time.Duration) (string, error) {

	t := jwt.New(jwt.SigningMethodHS256)

	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiteAt)),
		},
		UserID: userID,
	}
	return t.SignedString([]byte(s.config.SignKey))
}
