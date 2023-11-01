package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"gocron.com/m/pkg/cron/cron_data"
	"gorm.io/gorm"
	"time"
)

func StartCron(db *gorm.DB) {
	s := gocron.NewScheduler(time.UTC)

	userLogCron := cron_data.NewUserLogCronImpl(db)

	fmt.Println("Cron is running")

	s.Every(10).Seconds().Do(func() {
		fmt.Println("cron is running every 10 seconds")
		register, err := userLogCron.GetUserLog()
		if err != nil {
			return
		}

		for _, data := range register {
			dataUser, err := userLogCron.StoreUserLogRegister(data)
			fmt.Println("User with name : ", data.Name, " has been add to user log with uuid : ", dataUser.Uuid)
			if err != nil {
				return
			}
		}
	})

	s.StartAsync()

}
