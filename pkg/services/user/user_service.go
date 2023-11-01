package user

import "gocron.com/m/pkg/dto/user"

type UserService interface {
	RegisterUser(request user.RegisterUserDTO) error
}
