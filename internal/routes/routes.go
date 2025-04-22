package routes

import (
	"github.com/guilhermetk/summarize/internal/handlers"
	"github.com/guilhermetk/summarize/internal/providers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.Static("/", "public")

	geminiProvider := providers.GeminiProvider{}
	summarizeHandler := handlers.SummarizeHandler{
		Provider: &geminiProvider,
	}

	e.GET("/summarize", summarizeHandler.HandleGetSummarize)
}
