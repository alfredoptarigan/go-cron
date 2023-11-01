package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gocron.com/m/pkg/dto/user"
	user2 "gocron.com/m/pkg/repositories/user"
)

type userServiceImpl struct {
	userRepository user2.UserRepository
}

var validate = validator.New()

func (u *userServiceImpl) RegisterUser(request user.RegisterUserDTO) error {
	if err := validate.Struct(request); err != nil {
		return err
	}

	err := u.userRepository.Register(request)

	if err != nil {
		fmt.Errorf("Error when register user: %v", err)
	}

	return nil
}

func NewUserService(userRepository user2.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}
