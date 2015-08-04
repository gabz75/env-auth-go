package core

import (
  "testing"

  "github.com/gabz75/auth-api/core"
  "github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
  token := core.GenerateToken()
  assert.NotNil(t, token)
  assert.Equal(t, len(token), 233)
}
