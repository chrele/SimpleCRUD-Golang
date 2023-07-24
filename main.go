package main

import (
	"fmt"
	"log"
	"simple-crud-golang/config"
	"simple-crud-golang/controller"
	"simple-crud-golang/model"
	"simple-crud-golang/repository"
	"simple-crud-golang/router"
	"simple-crud-golang/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("Run service . . .")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables (.env files)", err)
	}

	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	noteRepository := repository.NewNoteRepositoryImpl(db)

	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	noteController := controller.NewNoteController(noteService)

	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8000"))
}
