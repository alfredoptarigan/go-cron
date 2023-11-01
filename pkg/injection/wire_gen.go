// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/google/wire"
	"gocron.com/m/pkg/config"
	"gocron.com/m/pkg/controllers/user"
	user2 "gocron.com/m/pkg/repositories/user"
	user3 "gocron.com/m/pkg/services/user"
)

// Injectors from injector.go:

func InitializeUserController() user.UserController {
	db := config.InitDatabase()
	userRepository := user2.NewUserRepository(db)
	userService := user3.NewUserService(userRepository)
	userController := user.NewUserController(userService)
	return userController
}

// injector.go:

var initDatabase = wire.NewSet(config.InitDatabase)
