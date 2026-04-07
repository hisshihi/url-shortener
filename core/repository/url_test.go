package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/hisshihi/url-shortener/core/repository/mocks"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

func Test_urlRepository_Create(t *testing.T) {
	type mockBehavior func(m *mocks.MockURLDb, longURL, alias string)

	tests := []struct {
		name         string
		url          string
		alias        string
		mockBehavior mockBehavior
		want         string
		wantErr      bool
	}{
		{
			name:  "success",
			url:   "https://google.com/long/path",
			alias: "abc12345",
			mockBehavior: func(m *mocks.MockURLDb, longURL, alias string) {
				m.EXPECT().
					Exec(mock.Anything, "INSERT INTO urls (long_url, alias) VALUES ($1, $2)", []interface{}{longURL, alias}).
					Return(pgconn.CommandTag{}, nil)
			},
			want:    "abc12345",
			wantErr: false,
		},
		{
			name: "fail",
			url:  "https://google.com/long/path",
			mockBehavior: func(m *mocks.MockURLDb, longURL, alias string) {
				m.EXPECT().
					Exec(mock.Anything, mock.AnythingOfType("string"), []interface{}{longURL, alias}).
					Return(pgconn.CommandTag{}, errors.New("db connection timeout"))
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := mocks.NewMockURLDb(t)

			tt.mockBehavior(db, tt.url, tt.alias)

			r := NewURLRepository(db)

			got, err := r.Create(context.Background(), tt.url, tt.alias)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
