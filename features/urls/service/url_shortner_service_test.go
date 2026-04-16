package service

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/hisshihi/url-shortener/features/urls/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_urlService_CreateShortURL_success(t *testing.T) {
	const inputURL = "https://google.com/long/path"
	const prefix = "http://shortner/"

	t.Run("success", func(t *testing.T) {
		mockRepo := mocks.NewMockURLRepo(t)
		mockRepo.
			EXPECT().
			Create(mock.Anything, inputURL, mock.MatchedBy(func(alias string) bool {
				if !strings.HasPrefix(alias, prefix) {
					return false
				}
				suffix := strings.TrimPrefix(alias, prefix)
				return len(suffix) == 8
			})).
			// В реальном репозитории возвращается alias, который ему передали.
			RunAndReturn(func(_ context.Context, _ string, alias string) (string, error) {
				return alias, nil
			}).
			Once()

		s := &URLService{
			urlRepo: mockRepo,
		}

		got, err := s.CreateShortURL(context.Background(), inputURL)
		assert.NoError(t, err)
		assert.True(t, strings.HasPrefix(got, prefix))
		assert.Len(t, strings.TrimPrefix(got, prefix), 8)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo := mocks.NewMockURLRepo(t)
		mockRepo.
			EXPECT().
			Create(mock.Anything, inputURL, mock.MatchedBy(func(alias string) bool {
				if !strings.HasPrefix(alias, prefix) {
					return false
				}
				suffix := strings.TrimPrefix(alias, prefix)
				return len(suffix) == 8
			})).
			RunAndReturn(func(_ context.Context, _ string, _ string) (string, error) {
				return "", errors.New("repo error")
			}).
			Once()

		s := &URLService{
			urlRepo: mockRepo,
		}
		_, err := s.CreateShortURL(context.Background(), inputURL)
		assert.Error(t, err)
	})

	t.Run("not valid url", func(t *testing.T) {
		url := "/path"
		mockRepo := mocks.NewMockURLRepo(t)
		s := &URLService{mockRepo}
		_, err := s.CreateShortURL(context.Background(), url)
		assert.ErrorIs(t, err, ErrInvalidURL)
	})
}

func Test_validateURL(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{rawURL: "https://google.com/long/path"},
			want: true,
		},
		{
			name: "invalid url",
			args: args{rawURL: "htt://google.com/long/path"},
			want: false,
		},
		{
			name: "empty url",
			args: args{rawURL: ""},
			want: false,
		},
		{
			name: "empty host",
			args: args{rawURL: "/abc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validateURL(tt.args.rawURL), "validateURL(%v)", tt.args.rawURL)
		})
	}
}
