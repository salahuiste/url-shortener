package handlers

import (
	"context"
	"net/http"

	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *services.URLService
}

func NewHandler(s *services.URLService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) Shorten(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}
	if c.BindJSON(&req) != nil {
		c.JSON(400, gin.H{"error": "invalid"})
		return
	}

	code, err := h.Service.Shorten(context.Background(), req.URL)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"short_url": "http://localhost:8080/" + code})
}

func (h *Handler) Resolve(c *gin.Context) {
	code := c.Param("code")

	url, err := h.Service.Resolve(context.Background(), code)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.Redirect(http.StatusFound, url)
}
