// Package db модели, функции и ошибки для работы с базой данных
package db

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrURLNotFound = gorm.ErrRecordNotFound
	ErrURLConflict = gorm.ErrInvalidData
	ErrURLExpired  = errors.New("срок действия короткой ссылки истек")
	ErrShortURLIsEmpty = errors.New("короткая ссылка не может быть пустой")
	ErrLongURLIsEmpty = errors.New("длинная ссылка не может быть пустой")
	ErrShortURLLengthInvalid = errors.New("короткая ссылка должна быть от 6 до 10 символов")
	ErrLongURLLengthInvalid = errors.New("длинная ссылка не может быть длиннее 2048 символов")
)

type URL struct {
	gorm.Model
	ShortURL  string `gorm:"uniqueIndex;not null"`
	LongURL   string `gorm:"not null"`
	ExpiresAt *time.Time
}

func (URL) TableName() string {
	return "urls"
}

// IsExpired проверяет, истек ли срок действия короткой ссылки
func (u *URL) IsExpired() bool {
	if u.ExpiresAt == nil {
		return false
	}
	return u.ExpiresAt.Before(time.Now())
}

// Validate валидирует URL
func (u *URL) Validate() error {
	if u.ShortURL == "" {
		return ErrShortURLIsEmpty
	}
	if len(u.ShortURL) < 6 || len(u.ShortURL) > 10 {
		return ErrShortURLLengthInvalid
	}
	if u.LongURL == "" {
		return ErrLongURLIsEmpty
	}
	if len(u.LongURL) > 2048 {
		return ErrLongURLLengthInvalid
	}
	return nil
}
