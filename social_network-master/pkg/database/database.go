package database

import (
	"Sunat/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase(configs *models.ConfigsModels) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		configs.Database.Host, configs.Database.User, configs.Database.Password,
		configs.Database.DbName, configs.Database.Port)
	database, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//addInformation(database)
	err = Migrations(database)
	if err != nil {
		return nil, err
	}

	return database, nil

}

func Migrations(db *gorm.DB) error {

	migrator := db.Migrator()
	log.Println("Everything is good ")
	err := migrator.AutoMigrate(&models.Users{}, &models.Tokens{}, &models.Posts{},
		&models.Comments{}, &models.Following{}, &models.Followers{}, &models.Likes{})

	if err != nil {
		log.Println(err)
		return err
	}
	//db.Exec(`alter table tokens
	//alter column expire set default current_timestamp + interval '3 hour';`)

	return nil
}
