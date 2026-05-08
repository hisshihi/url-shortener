// Package service предоставляет сервисы для работы с URL-shortner.
package service

import (
	"context"
	"errors"
	"log/slog"
	"net/url"

	"github.com/hisshihi/url-shortener/internal/pkg/util"
)

var (
	// ErrInvalidURL возвращает ошибку если URL недействителен
	ErrInvalidURL = errors.New("URL is invalid")
)

type URLRepository interface {
	Create(ctx context.Context, longURL, alias string) (string, error)
	SelectByAlias(ctx context.Context, alias string) (string, error)
}

type UrlShorterServiceCreate interface {
	CreateShortURL(ctx context.Context, rawUrl string) (string, error)
}

type UrlShorterServiceSelect interface {
	SelectByAlias(ctx context.Context, alias string) (string, error)
}

type UrlShorterService interface {
	UrlShorterServiceCreate
	UrlShorterServiceSelect
}

type urlShorter struct {
	urlRepo URLRepository
}

func NewURLService(urlRepository URLRepository) UrlShorterService {
	slog.Info("сервис url создан")
	return &urlShorter{
		urlRepo: urlRepository,
	}
}

// CreateShortURL создает короткий URL из длинного
func (s *urlShorter) CreateShortURL(ctx context.Context, rawUrl string) (string, error) {
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

func (s *urlShorter) SelectByAlias(ctx context.Context, alias string) (string, error) {
	isValid := validateURL(alias)
	if !isValid {
		return "", ErrInvalidURL
	}

	return s.urlRepo.SelectByAlias(ctx, alias)
}
