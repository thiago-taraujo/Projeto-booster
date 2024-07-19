package main

import (
	"Projeto-booster/configs"
	"Projeto-booster/internal/entity"
	"Projeto-booster/internal/infra/database"
	"Projeto-booster/internal/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"

	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Projeto Booster API Event
// @version 0.1
// @description An event API developed on Booster program
// @contact.name Thiago Araujo

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.Load(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("datasource.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Event{}, &entity.User{})

	eventDB := database.NewEvent(db)
	eventHandler := handlers.NewEventHandler(eventDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/events", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/{id}", eventHandler.GetEvent)
		r.Post("/", eventHandler.CreateEvent)
		r.Put("/{id}", eventHandler.UpdateEvent)
		r.Delete("/{id}", eventHandler.DeleteEvent)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("/docs/doc.json")))

	http.ListenAndServe(":8000", r)

}