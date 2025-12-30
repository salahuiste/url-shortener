package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"url-shortener/internal/repository"

	"github.com/go-redis/redis/v8"
)

type URLService struct {
	Repo  *repository.URLRepository
	Cache *redis.Client
}

func NewURLService(repo *repository.URLRepository, cache *redis.Client) *URLService {
	return &URLService{Repo: repo, Cache: cache}
}

// generate 6 char code
func generateCode() string {
	b := make([]byte, 4)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}

func (s *URLService) Shorten(ctx context.Context, url string) (string, error) {
	code := generateCode()
	err := s.Repo.Save(ctx, url, code)
	if err != nil {
		return "", err
	}

	// cache for faster future lookups
	s.Cache.Set(ctx, code, url, time.Hour*24)

	return code, nil
}

func (s *URLService) Resolve(ctx context.Context, code string) (string, error) {
	// Check Redis first
	url, err := s.Cache.Get(ctx, code).Result()
	if err == nil {
		return url, nil
	}

	// Fallback to DB
	url, err = s.Repo.Get(ctx, code)
	if err != nil {
		return "", err
	}

	// Store again in Redis
	s.Cache.Set(ctx, code, url, time.Hour*24)

	return url, nil
}
