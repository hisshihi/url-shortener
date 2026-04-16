// Package service предоставляет сервисы для работы с URL-shortner.
package service

import (
	"context"
	"errors"
	"log/slog"
	"net/url"

	"github.com/hisshihi/url-shortener/core/pkg/util"
)

var (
	// ErrInvalidURL возвращает ошибку если URL недействителен
	ErrInvalidURL = errors.New("URL is invalid")
)

type URLRepository interface {
	Create(ctx context.Context, longURL, alias string) (string, error)
	SelectByAlias(ctx context.Context, alias string) (string, error)
}

type URLService struct {
	urlRepo URLRepository
}

func NewURLService(urlRepository URLRepository) *URLService {
	slog.Info("сервис url создан")
	return &URLService{
		urlRepo: urlRepository,
	}
}

// CreateShortURL создает короткий URL из длинного
func (s *URLService) CreateShortURL(ctx context.Context, rawUrl string) (string, error) {
	isValid := validateURL(rawUrl)
	if !isValid {
		return "", ErrInvalidURL
	}

	alias, err := util.GenerateRandomString(8)
	if err != nil {
		return "", err
	}

	shortAlias := "http://shortner/" + alias
	return s.urlRepo.Create(ctx, rawUrl, shortAlias)
}

func validateURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if u.Host == "" {
		return false
	}

	if u.Scheme == "" || u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}

func (s *URLService) SelectByAlias(ctx context.Context, alias string) (string, error) {
	isValid := validateURL(alias)
	if !isValid {
		return "", ErrInvalidURL
	}

	return s.urlRepo.SelectByAlias(ctx, alias)
}
