package main

import (
	"log"

	"./common"
	"./persons"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&persons.Person{})
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	LoadEnv()
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	r.GET("/people/", persons.GetPersons)
	r.GET("/people/:id", persons.GetPerson)
	r.POST("/people", persons.CreatePerson)
	r.PUT("/people/:id", persons.UpdatePerson)
	r.DELETE("/peaple/:id", persons.DeletePerson)

	r.Run(":8080")
}
