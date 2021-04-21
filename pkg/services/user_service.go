package services

import (
	"github.com/wtfcode/hermes-api/pkg/helpers"
	"github.com/wtfcode/hermes-api/pkg/models"
	"gorm.io/gorm"
)

type IUserService interface {
	GetAllUsers(int) []models.User
	GetByID(int) (*models.User, error)
	GetByUsername(string) (*models.User, error)
	DeleteUser(int) error
	UpdateUser(models.User, uint) error
	CreateUser(models.User) error
}

type UserService struct {
	db *gorm.DB
}

const pageSize = 10

func NewUserService(db *gorm.DB) IUserService {
	return &UserService{db: db}
}

func (u *UserService) GetAllUsers(page int) []models.User {
	users := []models.User{}
	u.db.Scopes(helpers.Paginate(page, pageSize)).Select("first_name", "last_name", "email", "username").Find(&users)
	return users
}

func (u *UserService) GetByID(id int) (*models.User, error) {
	var user models.User
	if err := u.db.Select("first_name", "last_name", "email", "username").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := u.db.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) CreateUser(user models.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserService) DeleteUser(id int) error {
	return u.db.Delete(&models.User{}, id).Error
}

func (u *UserService) UpdateUser(user models.User, id uint) error {
	return u.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}
