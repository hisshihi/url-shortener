// Package service предоставляет сервисы для работы с URL-shortner.
package service

import (
	"context"
	"errors"
	"strings"

	"github.com/hisshihi/url-shortener/pkg/util"
)

var (
	// ErrInvalidURL возвращает ошибку если URL недействителен
	ErrInvalidURL = errors.New("URL is invalid")
)

// StringGenerator функция которая генерирует случайную строку заданной длины
type StringGenerator func(n int) (string, error)

// URLShortnerService сервис для создания коротких URL
type URLShortnerService struct {
	StringGenerator StringGenerator
}

// NewURLShortnerService создает новый URLShortnerService
func NewURLShortnerService() *URLShortnerService {
	return &URLShortnerService{
		StringGenerator: util.GenerateRandomString,
	}
}

// CreateShortURL создает короткий URL из длинного
func (s *URLShortnerService) CreateShortURL(ctx context.Context, url string) (string, error) {
	alias, err := s.StringGenerator(6)
	if err != nil {
		return "", err
	}

	if url == "" || len(url) <= 0 || (!strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") && !strings.HasPrefix(url, "ftp") && !strings.HasPrefix(url, "ftps") && !strings.Contains(url, "://")) {
		return "", ErrInvalidURL
	}

	domainFromURL := strings.Split(url, "/")[2]

	return "https://" + domainFromURL + "/" + alias, nil
}
