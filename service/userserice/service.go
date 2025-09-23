package userservice

import (
	"fmt"
	entity "suggestApp/enity"
	"suggestApp/pkg/phoneNumber"

	"golang.org/x/crypto/bcrypt"
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
	Password    string `json:"password"`
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

	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password is too short")
	}

	user := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    GetHash(req.Password),
	}

	createdUser, Err := s.repo.Register(user)
	if Err != nil {
		return RegisterResponse{}, fmt.Errorf("error registering user: %w", Err)
	}

	return RegisterResponse{createdUser}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	user entity.User
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	panic("")
}


func GetHash(data string ) string {
	hashedPassword := []byte(data)
	hashedPasswordString := string(hashedPassword)
	return hashedPasswordString
