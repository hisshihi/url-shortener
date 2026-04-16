// Package service предоставляет сервисы для работы с URL-shortner.
package service

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/hisshihi/url-shortener/core/pkg/util"
)

var (
	// ErrInvalidURL возвращает ошибку если URL недействителен
	ErrInvalidURL = errors.New("URL is invalid")
)

type URLRepository interface {
	Create(ctx context.Context, longURL, alias string) (string, error)
}

type URLService interface {
	CreateShortURL(ctx context.Context, longURL string) (string, error)
}

type urlService struct {
	urlRepo URLRepository
}

func NewURLService(urlRepository URLRepository) URLService {
	slog.Info("сервис url создан")
	return &urlService{
		urlRepo: urlRepository,
	}
}

// CreateShortURL создает короткий URL из длинного
func (s *urlService) CreateShortURL(ctx context.Context, url string) (string, error) {
	if !isValidURL(url) {
		return "", ErrInvalidURL
	}

	alias, err := util.GenerateRandomString(8)
	if err != nil {
		return "", err
	}

	domain := strings.Split(url, "/")[1]

	urlAlias := domain + "/" + alias

	shortURL, err := s.urlRepo.Create(ctx, url, urlAlias)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func isValidURL(url string) bool {
	if url == "" || len(url) <= 0 || (!strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") && !strings.HasPrefix(url, "ftp") && !strings.HasPrefix(url, "ftps") && !strings.Contains(url, "://")) {
		return false
	}
	return true
}
