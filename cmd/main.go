package main

import (
	"log"
	"url-shortener/internal/handlers"
	"url-shortener/internal/repository"
	"url-shortener/internal/services"
	"url-shortener/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.NewPostgres()
	if err != nil {
		log.Fatalf("fail to initialize database : %s", err.Error())
	}
	redis := pkg.NewRedis()

	repo := repository.NewURLRepository(db)
	service := services.NewURLService(repo, redis)
	h := handlers.NewHandler(service)

	r := gin.Default()
	r.POST("/shorten", h.Shorten)
	r.GET("/:code", h.Resolve)

	r.Run(":8080")
}
