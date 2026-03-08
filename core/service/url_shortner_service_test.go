package service_test

import (
	"errors"
	"testing"

	"github.com/hisshihi/url-shortener/core/service"
)

func TestURLShortnerService_CreateShortURL(t *testing.T) {
	tests := []struct {
		name      string
		URL       string
		generator service.StringGenerator
		want      string
		wantErr   bool
	}{
		{
			name: "correct URL",
			URL:  "https://example.com/test",
			generator: func(n int) (string, error) {
				return "abc123", nil
			},
			want:    "https://example.com/abc123",
			wantErr: false,
		},
		{
			name: "incurrect URL",
			URL:  "https://example.com/test",
			generator: func(n int) (string, error) {
				return "", service.ErrInvalidURL
			},
			wantErr: true,
		},
		{
			name: "generator error",
			URL:  "https://example.com/test",
			generator: func(n int) (string, error) {
				return "", errors.New("generator error")
			},
			wantErr: true,
		},
		{
			name:    "err URL",
			URL:     "invalid-url",
			generator: func(n int) (string, error) {
				return "abc123", nil
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewURLShortnerService()
			if tt.generator != nil {
				svc.StringGenerator = tt.generator
			}
			got, err := svc.CreateShortURL(t.Context(), tt.URL)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr=%v, got err=%v", tt.wantErr, err)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}

		})
	}
}
