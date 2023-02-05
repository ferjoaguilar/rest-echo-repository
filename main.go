package main

import (
	"context"
	"log"
	"os"

	"github.com/ferjoaguilar/rest-echo-repository/handlers"
	"github.com/ferjoaguilar/rest-echo-repository/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(bindRoutes)
}

func bindRoutes(s server.Sever, e *echo.Echo) {
	e.GET("/health-check", handlers.HealhCkeckHandler(s))
}
