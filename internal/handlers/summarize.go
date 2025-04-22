package handlers

import (
	"net/http"

	"github.com/guilhermetk/summarize/internal/types"
	"github.com/labstack/echo/v4"
)

type SummarizeHandler struct {
	Provider types.Provider
}

func (p *SummarizeHandler) HandleGetSummarize(c echo.Context) error {
	text := c.QueryParams().Get("text")
	summarized := p.Provider.Summarize(text)
	return c.String(http.StatusOK, summarized)
}
