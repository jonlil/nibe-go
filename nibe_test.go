package nibe

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "log"
)

func TestRefreshToken_should_refresh(t *testing.T) {
  rtr, err := RefreshToken("mysecretrefreshtoken")
  log.Println(rtr)
  assert.Equal(t, nil, err, "Err should be nil")
}
