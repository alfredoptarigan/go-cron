package user

import "gocron.com/m/pkg/dto/user"

type UserRepository interface {
	Register(request user.RegisterUserDTO) error
	AutoMigrate() error
}
