package gorm

import (
	"fmt"
	"github.com/franBarrientos/infrastructure/config"
	"github.com/franBarrientos/infrastructure/gorm/entities_db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDbConnection(env *config.Env) (*gorm.DB, error) {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var err error

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println("Error connecting to database: ", err)
	} else {
		log.Println("Connected to database")
		//syncronize tables...
		DB.AutoMigrate(entities_db.PersonalData{}, entities_db.User{}, entities_db.Organizer{}, entities_db.Place{}, entities_db.Event{})
	}
	return DB, err
}
