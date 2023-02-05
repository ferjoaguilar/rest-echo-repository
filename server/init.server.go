package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Port string `validate:"required"`
}

type Sever interface {
	Config() *Config
}

type broker struct {
	config       *Config
	echoInstance echo.Echo
}

func (b *broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*broker, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(config)
	if err != nil {
		return nil, err
	}

	broker := &broker{
		config:       config,
		echoInstance: *echo.New(),
	}

	return broker, nil
}

func (b *broker) Start(binder func(s Sever, e *echo.Echo)) {
	b.echoInstance = *echo.New()

	binder(b, &b.echoInstance)

	// Cors
	b.echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	//b.echoInstance.Logger.Fatal(b.echoInstance.Start(":" + b.config.Port))

	err := b.echoInstance.Start(":" + b.config.Port)
	if err != http.ErrServerClosed {
		log.Fatal("ListenAndServe: ", err)
	}
}
