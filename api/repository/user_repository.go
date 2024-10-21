package repository

import (
	"basic/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	FindByID(id uint) (models.User, error)
	Update(user models.User) error
	Delete(user models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAllUsers() ([]models.User, error) { // Implementasi fungsi untuk mendapatkan semua user
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) { // Implementasi fungsi untuk mendapatkan user berdasarkan ID
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *userRepository) Update(user models.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) Delete(user models.User) error {
	return r.db.Delete(&user).Error
}
