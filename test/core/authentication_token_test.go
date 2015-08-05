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

func TestExtractToken(t *testing.T) {
    validToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0MzkwNTYyNjJ9.Q5bhTFRUBJtctzvwTzLymM7Xy6L0_XCcfUwJFwoFZ2A8iQRjMTrCTl3pjHYWMiZAGylP_iimrRXgPCJa7Y2be-PtYgUOalZaTCuK2w3e9-zdzTuGJ8eDbXZUZbn-3LHDeUBpg18Il6aZWK9RJj1E50hTRUBOfBPdnTg_o3eYqAM"
    assert.Equal(t, core.ExtractToken("Bearer " + validToken), validToken)
    assert.Equal(t, core.ExtractToken(validToken), "")
    assert.Equal(t, core.ExtractToken("Bearer"), "")
}
