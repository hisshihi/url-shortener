package util_test

import (
	"testing"

	"github.com/hisshihi/url-shortener/pkg/util"
)

func TestGenerateRandomString(t *testing.T) {
	t.Run("correct length", func(t *testing.T) {
		got, err := util.GenerateRandomString(6)
		if err != nil {
			t.Fatalf("GenerateRandomString() failed: %v", err)
		}
		if len(got) != 6 {
			t.Errorf("GenerateRandomString() length = %d, want 6", len(got))
		}
	})
	t.Run("incorrect length", func(t *testing.T) {
		_, err := util.GenerateRandomString(0)
		if err == nil {
			t.Errorf("GenerateRandomString() should return an error for 0 length")
		}
	})
}
