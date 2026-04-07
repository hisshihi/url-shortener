package service

import (
	"context"
	"strings"
	"testing"

	"github.com/hisshihi/url-shortener/core/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_urlService_CreateShortURL(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		repoAlias string
		wantErr   error
	}{
		{
			name:      "success",
			inputURL:  "https://google.com/long/path",
			repoAlias: "abc12345",
			wantErr:   nil,
		},
		{
			name:     "not 8 characters",
			inputURL: "invalid url — no scheme",
			wantErr:  ErrInvalidURL,
		},
		{
			name:     "empty url",
			inputURL: "",
			wantErr:  ErrInvalidURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewMockURLRepo(t)

			if tt.wantErr == nil {
				mockRepo.On("Create", context.Background(), tt.inputURL, mock.AnythingOfType("string")).
					Once().
					Return(tt.repoAlias, nil)
			}

			s := &urlService{urlRepo: mockRepo}
			got, err := s.CreateShortURL(context.Background(), tt.inputURL)

			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr == nil {
				parts := strings.Split(got, "/")
				alias := parts[len(parts)-1]
				assert.Len(t, alias, 8)
				t.Logf("got: %s, alias: %s", got, alias)

				assert.Equal(t, tt.repoAlias, alias)
			}
		})
	}
}
