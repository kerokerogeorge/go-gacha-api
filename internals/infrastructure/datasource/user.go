package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) *userRepository {
	db := database
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(name string, token string, address string) (string, error) {
	user := model.User{Name: name, Token: token, Address: address}
	database := ur.db.Table("users").Create(&user)
	if database.Error != nil {
		return "", database.Error
	}
	return token, nil
}

func (ur *userRepository) GetUser(token string) (*model.User, error) {
	var user User
	err := ur.db.Table("users").Where("token = ?", token).First(&user).Error
	if err != nil {
		return nil, nil
	}
	return ur.ToUserModel(user), nil
}

func (ur *userRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, nil
	}
	return users, nil
}

func (ur *userRepository) UpdateUser(user *model.User, name string) (*model.User, error) {
	database := ur.db.Model(&user).Updates(name)
	if database.Error != nil {
		return nil, database.Error
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(user *model.User) error {
	err := ur.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) ToUserModel(user User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Token:     user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
