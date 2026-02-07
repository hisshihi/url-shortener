// Package db модели, функции и ошибки для работы с базой данных
package db

import (
	"strings"
	"testing"
	"time"
)

func TestURL_IsExpired(t *testing.T) {
	tests := []struct {
		name      string
		expiresAt *time.Time
		want      bool
	}{
		{
			name:      "не истекшая ссылка",
			expiresAt: nil,
			want:      false,
		},
		{
			name:      "истекшая ссылка",
			expiresAt: timePtr(time.Now().Add(-1 * time.Hour)),
			want:      true,
		},
		{
			name:      "не истекшая ссылка с будущей датой",
			expiresAt: timePtr(time.Now().Add(1 * time.Hour)),
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URL{
				ExpiresAt: tt.expiresAt,
			}
			if got := u.IsExpired(); got != tt.want {
				t.Errorf("URL.IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURL_Validate(t *testing.T) {
	tests := []struct {
		name    string
		url     *URL
		wantErr bool
	}{
		{
			name: "верный URL",
			url: &URL{
				ShortURL: "abc123",
				LongURL:  "https://example.com",
			},
			wantErr: false,
		},
		{
			name: "короткая ссылка пустая",
			url: &URL{
				ShortURL: "",
				LongURL:  "https://example.com",
			},
			wantErr: true,
		},
		{
			name: "длинная ссылка пустая",
			url: &URL{
				ShortURL: "abc123",
				LongURL:  "",
			},
			wantErr: true,
		},
		{
			name: "Короткая ссылка слишком короткая",
			url: &URL{
				ShortURL: "abc",
				LongURL:  "https://example.com",
			},
			wantErr: true,
		},
		{
			name: "Длинная ссылка слишком длинная",
			url: &URL{
				ShortURL: "abc123",
				LongURL:  strings.Repeat("a", 2049),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URL{
				tt.url.Model,
				tt.url.ShortURL,
				tt.url.LongURL,
				tt.url.ExpiresAt,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("URL.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func timePtr(t time.Time) *time.Time {
	return &t
}
