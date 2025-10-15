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
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(id uint) (entity.User, error)
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
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	fmt.Println(user, exist, err)

	fmt.Println(req.Password, GetHash(req.Password))
	if err != nil {
		return LoginResponse{}, fmt.Errorf("error getting user by phone number: %w", err)
	}

	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password is incorrect")
	}

	comp := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if comp != nil {
		return LoginResponse{}, fmt.Errorf("password is incorrect", comp)
	}

	return LoginResponse{user}, nil
}

func GetHash(data string) string {
	Password := []byte(data)
	hashedPass, err := bcrypt.GenerateFromPassword(Password, bcrypt.MinCost)
	if err != nil {
		return ""
	}
	hashedPasswordString := string(hashedPass)
	return hashedPasswordString
}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("Unexpected error getting user by id: %w", err)
	}

	return ProfileResponse{Name: user.Name}, nil
}
