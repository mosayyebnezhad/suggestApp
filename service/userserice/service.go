package userservice

import (
	"fmt"
	entity "suggestApp/enity"
	"suggestApp/pkg/phoneNumber"
)

type repository interface {
	IsUniquePhoneNumber(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	user entity.User
}

func NewService(repo repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO verify phone number
	fmt.Println(req)
	if !phoneNumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is invalid")
	}

	if isUnique, UErr := s.repo.IsUniquePhoneNumber(req.PhoneNumber); UErr != nil || !isUnique {
		if UErr != nil {
			return RegisterResponse{}, fmt.Errorf("error checking phone number uniqueness: %w", UErr)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name is too short")
	}

	user := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}

	createdUser, Err := s.repo.Register(user)
	if Err != nil {
		return RegisterResponse{}, fmt.Errorf("error registering user: %w", Err)
	}

	return RegisterResponse{createdUser}, nil
}
