package main

import (
	"crud/models"
	"crud/routes"
)

func main() {

    db := models.SetupDB()
    db.AutoMigrate(&models.Task{})

    r := routes.SetupRoutes(db)
    r.Run()
}