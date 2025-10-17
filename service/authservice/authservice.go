package authservice

import (
	"fmt"
	"strings"
	entity "suggestApp/enity"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	signKey               string
	accessExpiretionTime  time.Duration
	refreshExpiretionTime time.Duration
	accessSubject         string
	refreshSubject        string
}

func New(signKey, accessSubject, refreshSubject string, accessExpiretionTime, refreshExpiretionTime time.Duration) Service {

	return Service{
		signKey:               signKey,
		accessExpiretionTime:  accessExpiretionTime,
		refreshExpiretionTime: accessExpiretionTime,
		accessSubject:         accessSubject,
		refreshSubject:        refreshSubject,
	}

}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.accessSubject, s.accessExpiretionTime)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.refreshSubject, s.refreshExpiretionTime)
}

func (s Service) ParseToken(BearerToken string) (*Claims, error) {

	tokenStr := strings.Replace(BearerToken, "bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
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
	return t.SignedString([]byte(s.signKey))
}
