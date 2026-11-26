package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	if err := os.Setenv("JWT_SECRET", "test-secret"); err != nil {
		t.Error(err)
	}
	defer func() {
		err := os.Unsetenv("JWT_SECRET")
		if err != nil {
			t.Error(err)
		}
	}()
	err := InitSecret()
	var token, errorGenerate = GenerateToken("test-user-id")
	assert.Nil(t, err)
	assert.Nil(t, errorGenerate)
	assert.NotEqual(t, token, "")
}
