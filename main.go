package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go-send-email/app"
	"go-send-email/controller"
	"go-send-email/helper"
	"go-send-email/model/domain"
	"go-send-email/repository"
	"go-send-email/service"
	"net/http"
)

func main() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")
	err := viper.ReadInConfig()
	helper.PanicIfError(err)

	db := app.NewDB()
	validate := validator.New()
	err = db.AutoMigrate(&domain.User{})
	helper.PanicIfError(err)

	emailConfig := viper.GetString("email")
	configMail := domain.ConfigMail{
		SmtpHost: "smtp.gmail.com",
		SmtpPort: 587,
		Name:     fmt.Sprintf("%s <%s>", viper.GetString("name"), emailConfig),
		Email:    emailConfig,
		Password: viper.GetString("password"),
	}
	mailService := service.NewMailServiceImpl(configMail)

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, mailService, validate)
	userController := controller.NewUserControllerImpl(userService)

	err = db.AutoMigrate(&domain.User{})
	helper.PanicIfError(err)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/users", func(router chi.Router) {
		router.Post("/", userController.Create)
	})

	server := http.Server{
		Handler: router,
		Addr:    ":3000",
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
