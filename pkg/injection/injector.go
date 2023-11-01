//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"gocron.com/m/pkg/config"
	user3 "gocron.com/m/pkg/controllers/user"
	user2 "gocron.com/m/pkg/repositories/user"
	"gocron.com/m/pkg/services/user"
)

var initDatabase = wire.NewSet(
	config.InitDatabase,
)

func InitializeUserController() user3.UserController {
	wire.Build(
		initDatabase,
		user3.NewUserController,
		user.NewUserService,
		user2.NewUserRepository,
	)

	return nil
}
