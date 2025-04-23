package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/guilhermetk/summarize/internal/config"
	"github.com/guilhermetk/summarize/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadEnv()

	s := echo.New()
	s.Use(middleware.Logger())
	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	routes.SetupRoutes(s)

	if err := s.Start(":3000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}

}
