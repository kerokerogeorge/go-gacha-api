package datasource

import (
	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) *userRepository {
	db := database
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(name string, token string) (string, error) {
	user := model.User{Name: name, Token: token}
	database := ur.db.Table("users").Create(&user)
	if database.Error != nil {
		return "", database.Error
	}
	return token, nil
}

func (ur *userRepository) GetUser(token string) (string, error) {
	var user model.User
	err := ur.db.Table("users").Where("token = ?", token).First(&user).Error
	if err != nil {
		return "", nil
	}
	return user.Name, nil
}
