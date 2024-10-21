package service

import (
	"basic/api/dto"
	"basic/models"
	"basic/api/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(dto dto.CreateUserRequest) (dto.CreateUserResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	// Prepare user model
	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to repository
	err = s.repository.Create(&user)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	// Prepare response
	response := dto.CreateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}
