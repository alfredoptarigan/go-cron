package user

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gocron.com/m/pkg/dto/user"
	"gocron.com/m/pkg/models"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func (u userRepositoryImpl) AutoMigrate() error {
	err := u.db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}

func (u userRepositoryImpl) Register(request user.RegisterUserDTO) error {
	var userModel models.User

	fmt.Println("test")

	userModel.Uuid = uuid.New().String()
	userModel.Name = request.Name
	userModel.Email = request.Email
	userModel.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))

	err := u.db.Create(&userModel).Error

	fmt.Println("userModel", userModel)
	if err != nil {
		logrus.Error("Error when create user", err)
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}
