package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mahmudinashar/go/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Shirohige",
		Email:    "1@gmail.com",
		Password: "password",
		Role:     1,
	},
	models.User{
		Nickname: "Blackbeard",
		Email:    "2@gmail.com",
		Password: "password",
		Role:     2,
	},
	models.User{
		Nickname: "Big Mom",
		Email:    "3@gmail.com",
		Password: "password",
		Role:     2,
	},
	models.User{
		Nickname: "Kaido",
		Email:    "4@gmail.com",
		Password: "password",
		Role:     2,
	},
	models.User{
		Nickname: "Shanks",
		Email:    "5@gmail.com",
		Password: "password",
		Role:     2,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
